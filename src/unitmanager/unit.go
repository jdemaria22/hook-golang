package unitmanager

import (
	"fmt"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"sync"
)

const (
	INFO_SIZE int = 13224
)

func info(address int, deep bool, gameUnit GameUnit) (GameUnit, error) {
	data, err := memory.ReadInt(HOOK.Process, int(address))
	if err != nil {
		fmt.Println("error in info. data: ", err)
		return gameUnit, err
	}

	dataBuff, err := memory.Read(HOOK.Process, data, INFO_SIZE)
	if err != nil {
		fmt.Println("error in info. dataBuff: ", err)
		return gameUnit, err
	}

	var wg sync.WaitGroup
	var off int
	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJPOS
		gameUnit.Position.X = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJPOS + 0x4
		gameUnit.Position.Y = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJPOS + 0x8
		gameUnit.Position.Z = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJDIRECTION
		gameUnit.Direction.X = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJDIRECTION + 0x4
		gameUnit.Direction.Y = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJDIRECTION + 0x8
		gameUnit.Direction.Z = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJCRIT
		gameUnit.Crit = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJEXPIRY
		gameUnit.Duration = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJMAGICRES
		gameUnit.MagicResist = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJBONUSARMOR
		gameUnit.BonusArmor = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJARMOR
		gameUnit.Armor = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJBONUSATK
		gameUnit.BonusAttack = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJBASEATK
		gameUnit.BaseAttack = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJMAXHEALTH
		gameUnit.MaxHealth = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJHEALTH
		gameUnit.Health = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJHEALTH
		gameUnit.Health = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJMOVESPEED
		gameUnit.MovementSpeed = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJATKSPEEDMULTI
		gameUnit.AttackSpeedMulti = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJABILITYPOWER
		gameUnit.AbilityPower = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJMAGICRES
		gameUnit.MagicResist = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJCRITMULTI
		gameUnit.CritMulti = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJNETWORKID
		gameUnit.NetworkID = memory.Int32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		gameUnit.IsAlive = gameUnit.SpawnCount%2 == 0
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJSPAWNCOUNT
		gameUnit.SpawnCount = memory.Int32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJINDEX
		gameUnit.ObjectIndex = memory.Int32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJSIZEMULTIPLIER
		gameUnit.SizeMultiplier = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJLVL
		gameUnit.Level = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		off = offset.OBJATKRANGE
		gameUnit.AttackRange = memory.Float32frombytes(dataBuff[off : off+4])
	}()

	wg.Wait()

	return gameUnit, nil
}

func addChampInfoFromJson(gameUnit GameUnit) GameUnit {
	for i := 0; i < len(UNIT_DATA); i++ {
		if UNIT_DATA[i].Name == gameUnit.Name {
			gameUnit.AttackRangeJson = UNIT_DATA[i].AttackRange
			gameUnit.AcquisitionRange = UNIT_DATA[i].AcquisitionRange
			gameUnit.HealthBarHeight = UNIT_DATA[i].HealthBarHeight
			gameUnit.BaseMoveSpeed = UNIT_DATA[i].BaseMoveSpeed
			gameUnit.AttackSpeed = UNIT_DATA[i].AttackSpeed
			gameUnit.AttackSpeedRatio = UNIT_DATA[i].AttackSpeedRatio
			gameUnit.SelectionRadius = UNIT_DATA[i].SelectionRadius
			gameUnit.PathingRadius = UNIT_DATA[i].PathingRadius
			gameUnit.GameplayRadiusJson = UNIT_DATA[i].GameplayRadius
			gameUnit.BasicAtkMissileSpeed = UNIT_DATA[i].BasicAtkMissileSpeed
			gameUnit.BasicAtkWindup = UNIT_DATA[i].BasicAtkWindup
			gameUnit.Tags = UNIT_DATA[i].Tags
		}
		return gameUnit
	}
	return gameUnit
}
