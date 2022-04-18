package unitmanager

import (
	"fmt"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
)

type Buff struct {
	Name     string
	Count    int32
	CountAlt int32
	EndTime  float32
}

const (
	BUFF_SIZE             = 0x78
	ADDRESS_BUFF_ITERATOR = 0x8
)

func UpdateBuff(gameUnit GameUnit, data []byte) []Buff {
	var buffs []Buff
	buffArrayBgn := memory.Int32frombytes(data[offset.OBJBUFFMANAGER+offset.BUFFMANAGERENTRIESARRAY:])
	buffArrayEnd := memory.Int32frombytes(data[offset.OBJBUFFMANAGER+offset.BUFFMANAGERENDARRAY:])
	currentAddress := buffArrayBgn
	for currentAddress != buffArrayEnd {
		buffPointer, err := memory.ReadInt(HOOK.Process, int(currentAddress))
		if err != nil {
			fmt.Println("Error in buffPointer", err)
			currentAddress += ADDRESS_BUFF_ITERATOR
			continue
		}
		buffPointerAlloc, err := memory.Read(HOOK.Process, buffPointer, BUFF_SIZE)
		if err != nil {
			currentAddress += ADDRESS_BUFF_ITERATOR
			continue
		}

		dataBuff := memory.Int32frombytes(buffPointerAlloc[offset.BUFFNAME:])
		if dataBuff == 0 {
			currentAddress += ADDRESS_BUFF_ITERATOR
			continue
		}
		val, _ := memory.Read(HOOK.Process, int(dataBuff)-4+offset.BUFFNAME, 100)
		name := memory.CopyString(val)
		if name == "" {
			currentAddress += ADDRESS_BUFF_ITERATOR
			continue
		}
		var buff Buff
		buff.Name = name
		buff.Count = memory.Int32frombytes(buffPointerAlloc[offset.BUFFENTRYBUFFCOUNT:])
		buff.CountAlt = memory.Int32frombytes(buffPointerAlloc[offset.BUFFENTRYBUFFCOUNTALT:]) - memory.Int32frombytes(buffPointerAlloc[offset.BUFFENTRYBUFFCOUNTALT:])>>3
		buff.EndTime = memory.Float32frombytes(buffPointerAlloc[offset.BUFFENTRYBUFFENDTIME:])
		buffs = append(buffs, buff)
		currentAddress += ADDRESS_BUFF_ITERATOR
	}
	return buffs
}
