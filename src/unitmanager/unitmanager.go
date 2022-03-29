package unitmanager

type GamePosition struct {
	X float32
	Y float32
	Z float32
}

type UnitManager struct {
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

// func Update(hook hook.ProcessHook) (UnitManager, error) {
// 	var unitManager UnitManager
// 	objectManagerOffset, err := memory.ReadInt(hook, hook.ModuleBaseAddr+offset.OBJECTMANAGER)
// 	if err != nil {
// 		return unitManager, err
// 	}

// 	if objectManagerOffset <= 0 {
// 		return unitManager, fmt.Errorf("error to find objectManagerOffset")
// 	}

// }
