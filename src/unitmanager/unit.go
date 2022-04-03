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

func info(address int, deep bool) (GameUnit, error) {
	var gameUnit GameUnit
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

	// fmt.Println("gameUnit.Position ,", gameUnit.Position)

	// fmt.Println("gameUnit.Direction ,", gameUnit.Direction)

	// fmt.Println("gameUnit.Health ", gameUnit.Health)

	// fmt.Println("gameUnit.MaxHealth ", gameUnit.MaxHealth)

	// fmt.Println("gameUnit.BaseAttack ", gameUnit.BaseAttack)

	// fmt.Println("gameUnit.BonusAttack ", gameUnit.BonusAttack)

	// fmt.Println("gameUnit.Armor ", gameUnit.Armor)

	// fmt.Println("gameUnit.BonusArmor ", gameUnit.BonusArmor)

	// fmt.Println("gameUnit.MagicResist ", gameUnit.MagicResist)

	// fmt.Println("gameUnit.Duration ", gameUnit.Duration)

	// fmt.Println("gameUnit.Crit ", gameUnit.Crit)

	// fmt.Println("gameUnit.CritMulti ", gameUnit.CritMulti)

	// fmt.Println("gameUnit.MagicResist ", gameUnit.MagicResist)

	// fmt.Println("gameUnit.AbilityPower ", gameUnit.AbilityPower)

	// fmt.Println("gameUnit.AttackSpeedMulti ", gameUnit.AttackSpeedMulti)

	// fmt.Println("gameUnit.MovementSpeed ", gameUnit.MovementSpeed)

	// fmt.Println("gameUnit.AttackRange ", gameUnit.AttackRange)

	// fmt.Println("gameUnit.Level ", gameUnit.Level)

	// fmt.Println("gameUnit.SizeMultiplier ", gameUnit.SizeMultiplier)

	// fmt.Println("gameUnit.ObjectIndex ", gameUnit.ObjectIndex)

	// fmt.Println("gameUnit.SpawnCount ", gameUnit.SpawnCount)

	// fmt.Println("gameUnit.NetworkID ", gameUnit.NetworkID)
	// // gameUnit.IsTargetable = destfloat[0]

	// // gameUnit.IsVisible = destfloat[0]

	return gameUnit, nil
}
