package memory

import (
	"encoding/binary"
	"framework-memory-go/src/hook"
	"framework-memory-go/src/size"
	"framework-memory-go/src/win"
	"math"
	"unicode/utf16"
)

const (
	FILE_DEVICE_UNKNOWN uint32 = 0x00000022
	FILE_VALUE          uint32 = 0x6211
	METHOD_BUFFERED     uint32 = 0
	FILE_SPECIAL_ACCESS uint32 = 0
)

type KERNEL_READ_REQUEST struct {
	ProcessId win.HANDLE //target process id
	Address   int        // address of memory to start reading from
	pBuff     []byte     // return value
	Size      uint       // size of memory to read
}

func StringToUTF16Ptr(str string) *uint16 {
	wchars := utf16.Encode([]rune(str + "\x00"))
	return &wchars[0]
}

func ReadFloat(hook hook.ProcessHook, offsets int) (float32, error) {
	var value float32
	data, err := win.ReadProcessMemory(hook.Process, uint32(offsets), size.Float)
	if err != nil {
		return value, err
	}
	bits := binary.LittleEndian.Uint32(data)
	return math.Float32frombits(bits), nil
}

func ReadInt(hook hook.ProcessHook, offsets int) (int, error) {
	var value int
	data, err := win.ReadProcessMemory(hook.Process, uint32(offsets), size.Int)
	if err != nil {
		return value, err
	}
	bits := binary.LittleEndian.Uint32(data)
	return int(bits), nil
}

func Read(hook hook.ProcessHook, offsets int, size int) ([]byte, error) {
	data, err := win.ReadProcessMemory(hook.Process, uint32(offsets), uint(size))
	if err != nil {
		return data, err
	}
	return data, nil
}

func MemCopy(src []byte, dest int) {

}

func ctl_code(device_type, function, method, access uint32) uint32 {
	return (device_type << 16) | (access << 14) | (function << 2) | method
}
