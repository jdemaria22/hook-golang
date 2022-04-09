package unitmanager

import (
	"fmt"
	Hook "framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"sync"
)

type GamePosition struct {
	X float32
	Y float32
	Z float32
}

type GameUnit struct {
	Address           uint
	Name              string
	LastVisibleTime   float32
	Team              int32
	Health            float32
	MaxHealth         float32
	BaseAttack        float32
	BonusAttack       float32
	Armor             float32
	BonusArmor        float32
	MagicResist       float32
	Duration          float32
	IsVisible         bool
	ObjectIndex       int32
	Crit              float32
	CritMulti         float32
	AbilityPower      float32
	AttackSpeedMulti  float32
	MovementSpeed     float32
	NetworkID         int32
	SpawnCount        int32
	IsAlive           bool
	AttackRange       float32
	IsTargetable      bool
	Level             float32
	GameplayRadius    float32
	SizeMultiplier    float32
	IsChampion        bool
	IsImportantJungle bool
	Position          GamePosition
	Direction         GamePosition
}

type UnitManager struct {
	Champions []GameUnit
	Monsters  []GameUnit
	Minions   []GameUnit
	Turrets   []GameUnit
	Units     map[int32]GameUnit
}

const (
	OBJECT_MANAGER_BUFF int = 256
	MAX_UNITS           int = 500
	ARRAY_HERO_LIST     int = 0x04
	ARRAY_HERO_LENGTH   int = 0x08
)

var (
	HOOK        Hook.ProcessHook = Hook.HOOK
	unitReads   int              = 0
	UNITMANAGER UnitManager
	mu          = &sync.Mutex{}
)

func Update() error {
	var wg sync.WaitGroup
	var unitManager UnitManager
	UNITMANAGER = unitManager
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
	return nil
}

func updateChampions() {
	hero, err := memory.ReadInt(HOOK.Process, HOOK.ModuleBaseAddr+offset.AIHeroClient)
	if err != nil {
		fmt.Println("Error in AIHeroClient ", err)
	}
	heroArray, err := memory.ReadInt(HOOK.Process, hero+0x04)
	if err != nil {
		fmt.Println("Error in heroArray ", err)
	}
	heroArrayLen, err := memory.ReadInt(HOOK.Process, hero+0x08)
	if err != nil {
		fmt.Println("Error in heroArrayLen ", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < heroArrayLen*4; i += 4 {
		var gameUnit GameUnit
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			gameUnit, err = info(heroArray+i, true)
			if err != nil {
				fmt.Println("Error in updateChampions.info ", err)
			}
			mu.Lock()
			UNITMANAGER.Champions = append(UNITMANAGER.Champions, gameUnit)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
}

func updateMinions() {
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
	var wg sync.WaitGroup
	for i := 0; i < minionArrayLen*4; i += 4 {
		var gameUnit GameUnit
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			gameUnit, err = info(minionArray+i, true)
			if err != nil {
				fmt.Println("Error in updateMinions.info ", err)
			}
			mu.Lock()
			UNITMANAGER.Minions = append(UNITMANAGER.Minions, gameUnit)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
}
