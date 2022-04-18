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
	Icon                 *ebiten.Image
	Buffs                []Buff
	Spells               [6]Spell
}

type UnitManager struct {
	Champions map[int]GameUnit
	Minions   map[int]GameUnit
}

const (
	OBJECT_MANAGER_BUFF int = 256
	MAX_UNITS           int = 500
	ARRAY_HERO_LIST     int = 0x04
	ARRAY_HERO_LENGTH   int = 0x08
)

func init() {
	UNITMANAGER.Champions = make(map[int]GameUnit)
	UNITMANAGER.Minions = make(map[int]GameUnit)
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
	LOCALPLAYER  GameUnit
)

func Update() error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		updateChampions()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		updateMinions()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		updateMe()
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

func updateMinions() {
	minions := make(map[int]GameUnit)

	hero, err := memory.ReadInt(HOOK.Process, HOOK.ModuleBaseAddr+offset.AIMinionClient)
	if err != nil {
		fmt.Println("Error in AIMinionClient ", err)
	}
	if err != nil {
		fmt.Println(err)
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
		if val, ok := UNITMANAGER.Minions[idunit]; ok {
			gameUnit, err := infoMinion(idunit, false, val)
			if err != nil {
				fmt.Println("Error in updateChampions.info ", err)
			}
			minions[idunit] = gameUnit
		} else {
			var gameUnit GameUnit
			gameUnit, err = infoMinion(minionArray+i, true, gameUnit)
			if err != nil {
				fmt.Println("Error in updateMinions.info ", err)
			}
			minions[idunit] = gameUnit
		}
	}
	UNITMANAGER.Minions = minions
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
