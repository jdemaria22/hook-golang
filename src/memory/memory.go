package memory

import (
	"encoding/binary"
	"framework-memory-go/src/size"
	"framework-memory-go/src/win"
	"math"
	"strings"
	"unicode/utf16"
	"unsafe"
)

const (
	FILE_DEVICE_UNKNOWN uint32 = 0x00000022
	FILE_VALUE          uint32 = 0x6211
	METHOD_BUFFERED     uint32 = 0
	FILE_SPECIAL_ACCESS uint32 = 0
	BYTE_MAX_VALUE      int    = 127
)

type KERNEL_READ_REQUEST struct {
	ProcessId win.HANDLE
	Address   int
	pBuff     []byte
	Size      uint
}

func StringToUTF16Ptr(str string) *uint16 {
	wchars := utf16.Encode([]rune(str + "\x00"))
	return &wchars[0]
}

func ReadFloat(process win.HANDLE, offsets int) (float32, error) {
	var value float32
	data, err := win.ReadProcessMemory(process, uint32(offsets), size.Float)
	if err != nil {
		return value, err
	}
	bits := binary.LittleEndian.Uint32(data)
	return math.Float32frombits(bits), nil
}

func ReadInt(process win.HANDLE, offsets int) (int, error) {
	var value int
	data, err := win.ReadProcessMemory(process, uint32(offsets), size.Int)
	if err != nil {
		return value, err
	}
	bits := binary.LittleEndian.Uint32(data)
	return int(bits), nil
}

func Read(process win.HANDLE, offsets int, size int) ([]byte, error) {
	data, err := win.ReadProcessMemory(process, uint32(offsets), uint(size))
	if err != nil {
		return data, err
	}
	return data, nil
}

func Float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func CopyString(bytes []byte) string {
	var str strings.Builder
	for _, b := range bytes {
		c := int(b) & 0xFF
		if c == 0 {
			break
		}
		if c > BYTE_MAX_VALUE {
			return ""
		}
		str.WriteByte(b)
	}
	return str.String()
}

func Int32frombytes(bytes []byte) int32 {
	return int32(binary.LittleEndian.Uint32(bytes))
}

func CopyInt(data []byte, offset int) int32 {
	var destint int32
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destint)), unsafe.Sizeof(data)), data[offset:])
	return destint
}

func CopyFloat(data []byte, offset int) float32 {
	var destfloat float32
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destfloat)), unsafe.Sizeof(data)), data[offset:])
	return destfloat
}

func ctl_code(device_type, function, method, access uint32) uint32 {
	return (device_type << 16) | (access << 14) | (function << 2) | method
}
