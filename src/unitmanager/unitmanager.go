package unitmanager

import (
	"fmt"
	"framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"unsafe"
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
	Team              int
	Health            float32
	MaxHealth         float32
	BaseAttack        float32
	BonusAttack       float32
	Armor             float32
	BonusArmor        float32
	MagicResist       float32
	Duration          float32
	IsVisible         bool
	ObjectIndex       int
	Crit              float32
	CritMulti         float32
	AbilityPower      float32
	AttackSpeedMulti  float32
	MovementSpeed     float32
	NetworkID         int
	SpawnCount        int
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
	MAX_UNITS           int = 2048
)

func Update(hook hook.ProcessHook) (UnitManager, error) {
	var unitManager UnitManager
	objectManagerOffset, err := memory.ReadInt(hook, hook.ModuleBaseAddr+offset.OBJECTMANAGER)
	if err != nil {
		return unitManager, err
	}
	if objectManagerOffset <= 0 {
		return unitManager, fmt.Errorf("error to find objectManagerOffset")
	}

	objetManager, err := memory.Read(hook, objectManagerOffset, OBJECT_MANAGER_BUFF)
	if err != nil {
		return unitManager, err
	}

	scan := scanUnits(hook, objetManager, unitManager)
	if scan {
		return unitManager, nil
	}
	return unitManager, fmt.Errorf("Error in scan units.")
}

func scanUnits(hook hook.ProcessHook, objectMangaer []byte, unitManager UnitManager) bool {
	destint := make([]int32, 1)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(objectMangaer)), objectMangaer[offset.OBJECTMAPCOUNT:])
	if destint[0] <= 0 {
		return false
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(objectMangaer)), objectMangaer[offset.OBJECTMAPROOT:])
	rootUnitAddress := destint[0]
	if rootUnitAddress <= 0 {
		return false
	}

	var unitReads int32 = 0
	scanUnit(hook, rootUnitAddress, unitReads, unitManager)
	////
	// for unitsRead <= MAX_UNITS || notContains(nodes, rootUnitAddress) {
	// 	rootBuff, err := memory.Read(hook, int(rootUnitAddress), 0x18)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	unitsRead++
	// 	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[offset.OBJECTMAPNODENETID:])
	// 	networkId := destint[0]
	// 	if int(networkId) >= 0x40000000 {
	// 		copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[offset.OBJECTMAPNODEOBJECT:])
	// 		unitAddress := destint[0]
	// 	}
	// 	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[0x0:])
	// 	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[0x4:])
	// 	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[0x8:])
	// }

	return true
}

func scanUnit(hook hook.ProcessHook, address int32, unitReads int32, unitManager UnitManager) {
	destint := make([]int32, 1)
	nodes := make([]int32, MAX_UNITS)
	unitReads++
	if int(unitReads) >= MAX_UNITS || contains(nodes, address) {
		return
	}

	_ = append(nodes, address)
	rootBuff, err := memory.Read(hook, int(address), 0x30)
	if err != nil {
		return
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[offset.OBJECTMAPNODENETID:])
	networkId := destint[0]
	if int(networkId) >= 0x40000000 {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[offset.OBJECTMAPNODEOBJECT:])
		unitAddress := destint[0]
		updateUnit(hook, networkId, unitAddress, unitManager)
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[0x0:])
	scanUnit(hook, destint[0], unitReads, unitManager)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[0x4:])
	scanUnit(hook, destint[0], unitReads, unitManager)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint[0])), unsafe.Sizeof(rootBuff)), rootBuff[0x8:])
	scanUnit(hook, destint[0], unitReads, unitManager)
}

func updateUnit(hook hook.ProcessHook, networkId int32, address int32, unitManager UnitManager) {
	if address <= 0 {
		return
	}

	var unit GameUnit
	if val, ok := unitManager.Units[networkId]; ok {
		unit = val
		// update unit con deep en false y reemplazar el valor que se encuentra en el mapa de unidades
		unitManager.Units[networkId] = unit
	}
	// update unit con deep en true y agregarlo al mapa de la lista de unidades

	unitManager.Units[networkId] = unit

}

func contains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func notContains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return false
		}
	}
	return true
}
