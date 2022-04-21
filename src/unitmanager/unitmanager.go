package unitmanager

import (
	"fmt"
	Hook "framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type GamePosition struct {
	X float32
	Y float32
	Z float32
}

type GameUnit struct {
	Address              uint
	Name                 string
	LastVisibleTime      float32
	Team                 int
	Health               float32
	MaxHealth            float32
	BaseAttack           float32
	BonusAttack          float32
	Armor                float32
	BonusArmor           float32
	MagicResist          float32
	Duration             float32
	IsVisible            bool
	ObjectIndex          int32
	Crit                 float32
	CritMulti            float32
	AbilityPower         float32
	AttackSpeedMulti     float32
	MovementSpeed        float32
	NetworkID            int32
	SpawnCount           int32
	IsAlive              bool
	AttackRange          float32
	IsTargetable         bool
	Level                float32
	GameplayRadius       float32
	SizeMultiplier       float32
	IsChampion           bool
	IsImportantJungle    bool
	Position             GamePosition
	Direction            GamePosition
	HealthBarHeight      float32
	BaseMoveSpeed        float32
	AttackRangeJson      float32
	AttackSpeed          float32
	AttackSpeedRatio     float32
	AcquisitionRange     float32
	SelectionRadius      float32
	PathingRadius        float32
	GameplayRadiusJson   float32
	BasicAtkMissileSpeed float32
	BasicAtkWindup       float32
	Tags                 []string
	UnitType             int
	AiManager            AiManager
	Icon                 *ebiten.Image
	Buffs                []Buff
	Spells               [6]Spell
}

type UnitManager struct {
	Champions  map[int]GameUnit
	AllUnits   map[int]GameUnit
	Turrets    map[int]GameUnit
	Inhibitors map[int]GameUnit
}

const (
	OBJECT_MANAGER_BUFF int = 256
	MAX_UNITS           int = 500
	ARRAY_HERO_LIST     int = 0x04
	ARRAY_HERO_LENGTH   int = 0x08
)

func init() {
	UNITMANAGER.Champions = make(map[int]GameUnit)
	UNITMANAGER.AllUnits = make(map[int]GameUnit)
	UNITMANAGER.Turrets = make(map[int]GameUnit)
	UNITMANAGER.Inhibitors = make(map[int]GameUnit)
}

var (
	HOOK         Hook.ProcessHook = Hook.HOOK
	unitReads    int              = 0
	UNITMANAGER  UnitManager
	mu           = sync.RWMutex{}
	testMutex    sync.Mutex
	wg           sync.WaitGroup
	hero         = 0
	heroArray    = 0
	heroArrayLen = 0
	turret       = 0
	turretArray  = 0
	turArrayLen  = 0
	inhibitor    = 0
	inhibiArray  = 0
	inhArrayLen  = 0
	LOCALPLAYER  GameUnit
)

func Update() error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		updateChampions()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		updateMe()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		updateAllUnits()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		updateTurrets()
	}()
	wg.Wait()
	return nil
}

func updateChampions() {
	if hero == 0 {
		herovalue, err := memory.ReadInt(HOOK.Process, HOOK.ModuleBaseAddr+offset.AIHeroClient)
		if err != nil {
			fmt.Println("Error in AIHeroClient ", err)
		}
		hero = herovalue
	}

	if heroArray == 0 {
		heroArrayValue, err := memory.ReadInt(HOOK.Process, hero+0x04)
		if err != nil {
			fmt.Println("Error in heroArray ", err)
		}
		heroArray = heroArrayValue
	}

	if heroArrayLen == 0 {
		heroArrayLenValue, err := memory.ReadInt(HOOK.Process, hero+0x08)
		if err != nil {
			fmt.Println("Error in heroArrayLen ", err)
		}
		heroArrayLen = heroArrayLenValue
	}
	var err error
	for i := 0; i < heroArrayLen*4; i += 4 {
		idunit := heroArray + i
		if val, ok := UNITMANAGER.Champions[idunit]; ok {
			gameUnit, err := infoChampion(idunit, false, val)
			if err != nil {
				fmt.Println("Error in updateChampions.info ", err)
			}
			UNITMANAGER.Champions[idunit] = gameUnit
		} else {
			var gameUnit GameUnit
			gameUnit, err = infoChampion(idunit, true, gameUnit)
			if err != nil {
				fmt.Println("Error in updateChampions.info ", err)
			}
			UNITMANAGER.Champions[idunit] = gameUnit
		}
	}
}

func updateAllUnits() {
	allUnits := make(map[int]GameUnit)

	hero, err := memory.ReadInt(HOOK.Process, HOOK.ModuleBaseAddr+offset.AIMinionClient)
	if err != nil {
		fmt.Println("Error in AIMinionClient ", err)
	}
	minionArray, err := memory.ReadInt(HOOK.Process, hero+0x04)
	if err != nil {
		fmt.Println("Error in minionArray ", err)
	}
	minionArrayLen, err := memory.ReadInt(HOOK.Process, hero+0x08)
	if err != nil {
		fmt.Println("Error in minionArrayLen ", err)
	}
	for i := 0; i < minionArrayLen*4; i += 4 {
		idunit := minionArray + i
		if val, ok := UNITMANAGER.AllUnits[idunit]; ok {
			gameUnit, err := infoMinion(idunit, false, val)
			if err != nil {
				fmt.Println("Error in updateChampions.info ", err)
			}
			allUnits[idunit] = gameUnit
		} else {
			var gameUnit GameUnit
			gameUnit, err = infoMinion(idunit, true, gameUnit)
			if err != nil {
				fmt.Println("Error in updateAllUnits.info ", err)
			}
			allUnits[idunit] = gameUnit
		}
	}
	UNITMANAGER.AllUnits = allUnits
}

func updateTurrets() {
	if turret == 0 {
		turretvalue, err := memory.ReadInt(HOOK.Process, HOOK.ModuleBaseAddr+offset.AITurretClient)
		if err != nil {
			fmt.Println("Error in AITurretClient ", err)
		}
		turret = turretvalue
	}

	if turretArray == 0 {
		turretArrayValue, err := memory.ReadInt(HOOK.Process, turret+0x04)
		if err != nil {
			fmt.Println("Error in turretArray ", err)
		}
		turretArray = turretArrayValue
	}

	if turArrayLen == 0 {
		turretArrayLenValue, err := memory.ReadInt(HOOK.Process, turret+0x08)
		if err != nil {
			fmt.Println("Error in turretArrayLenValue ", err)
		}
		turArrayLen = turretArrayLenValue
	}
	var err error
	for i := 0; i < turArrayLen*4; i += 4 {
		idunit := turretArray + i
		if val, ok := UNITMANAGER.Turrets[idunit]; ok {
			gameUnit, err := infoTurret(idunit, false, val)
			if err != nil {
				fmt.Println("Error in updateTurrets.info ", err)
			}
			UNITMANAGER.Turrets[idunit] = gameUnit
		} else {
			var gameUnit GameUnit
			gameUnit, err = infoTurret(idunit, true, gameUnit)
			if err != nil {
				fmt.Println("Error in updateTurrets.info ", err)
			}
			UNITMANAGER.Turrets[idunit] = gameUnit
		}
	}
}

func updateInhibitors() {
	if inhibitor == 0 {
		inhibitorvalue, err := memory.ReadInt(HOOK.Process, HOOK.ModuleBaseAddr+offset.AIBuildingList)
		if err != nil {
			fmt.Println("Error in AIInhibitorList ", err)
		}
		inhibitor = inhibitorvalue
	}

	if inhibiArray == 0 {
		inhibitorArrayValue, err := memory.ReadInt(HOOK.Process, inhibitor+0x04)
		if err != nil {
			fmt.Println("Error in inhibiArray ", err)
		}
		inhibiArray = inhibitorArrayValue
	}

	if inhArrayLen == 0 {
		inhibitorArrayLenValue, err := memory.ReadInt(HOOK.Process, inhibitor+0x08)
		if err != nil {
			fmt.Println("Error in inhArrayLen ", err)
		}
		inhArrayLen = inhibitorArrayLenValue
	}
	var err error
	for i := 0; i < inhArrayLen*4; i += 4 {
		idunit := inhibiArray + i
		if val, ok := UNITMANAGER.Inhibitors[idunit]; ok {
			gameUnit, err := infoInhibitor(idunit, false, val)
			if err != nil {
				fmt.Println("Error in updateInhibitors.info ", err)
			}
			UNITMANAGER.Inhibitors[idunit] = gameUnit
		} else {
			var gameUnit GameUnit
			gameUnit, err = infoInhibitor(idunit, true, gameUnit)
			if err != nil {
				fmt.Println("Error in updateInhibitors.info ", err)
			}
			UNITMANAGER.Inhibitors[idunit] = gameUnit
		}
	}
}

func updateMe() {
	address, err := memory.ReadInt(HOOK.Process, HOOK.ModuleBaseAddr+offset.LOCALPLAYER)
	if err != nil {
		fmt.Println("Error in Address for LocalPlayer ", err)
	}

	networkID, err := memory.ReadInt(HOOK.Process, address+offset.OBJNETWORKID)
	if err != nil {
		fmt.Println("Error in networkID for LocalPlayer ", err)
	}
	for _, element := range UNITMANAGER.Champions {
		if element.NetworkID == int32(networkID) {
			LOCALPLAYER = element
		}
	}
}
