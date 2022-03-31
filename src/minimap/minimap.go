package minimap

import (
	"fmt"
	"framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"framework-memory-go/src/renderer"
	"unsafe"
)

type Minimap struct {
	Width  float32
	Height float32
	X      float32
	Y      float32
}

const (
	Y_ADD_VALUE           int     = 4
	WORLD_SCALE           float32 = 15000
	MINIMAP_HUD_BUFF_SIZE int     = 0x80
)

func Update(hook hook.ProcessHook) (Minimap, error) {
	var minimap Minimap
	minimapObject, err := memory.ReadInt(hook, hook.ModuleBaseAddr+offset.MINIMAPOBJECT)
	if err != nil {
		return minimap, err
	}

	if minimapObject <= 0 {
		return minimap, fmt.Errorf("error to find minimapObject")
	}

	minimapHUD, err := memory.ReadInt(hook, minimapObject+offset.MINIMAPOBJECTHUD)
	if err != nil {
		return minimap, err
	}

	if minimapHUD <= 0 {
		return minimap, fmt.Errorf("error to find minimapHUD")
	}

	minimapHUDBuff, err := memory.Read(hook, minimapHUD, MINIMAP_HUD_BUFF_SIZE)
	if err != nil {
		return minimap, err
	}

	destfloat := make([]float32, 1)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destfloat[0])), unsafe.Sizeof(minimapHUDBuff)), minimapHUDBuff[offset.MINIMAPHUDPOS:])
	minimap.X = destfloat[0]

	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destfloat[0])), unsafe.Sizeof(minimapHUDBuff)), minimapHUDBuff[offset.MINIMAPHUDPOS+Y_ADD_VALUE:])
	minimap.Y = destfloat[0]

	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destfloat[0])), unsafe.Sizeof(minimapHUDBuff)), minimapHUDBuff[offset.MINIMAPHUDSIZE:])
	minimap.Width = destfloat[0]

	copy(unsafe.Slice((*byte)(unsafe.Pointer(&destfloat[0])), unsafe.Sizeof(minimapHUDBuff)), minimapHUDBuff[offset.MINIMAPHUDSIZE+Y_ADD_VALUE:])
	minimap.Height = destfloat[0]

	return minimap, nil
}

func MinimapToScreen(x float32, y float32, z float32, minimap Minimap) renderer.ScreenPosition {
	var screenPosition renderer.ScreenPosition
	var rx = x / WORLD_SCALE
	var ry = z / WORLD_SCALE

	rx = minimap.X + rx*minimap.Width
	ry = minimap.Y + minimap.Height - (ry * minimap.Height)
	screenPosition.X = rx
	screenPosition.Y = ry
	return screenPosition
}
