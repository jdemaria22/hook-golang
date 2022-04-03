package minimap

import (
	"fmt"
	Hook "framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"framework-memory-go/src/renderer"
)

type Minimap struct {
	Width  float32
	Height float32
	X      float32
	Y      float32
}

const (
	Y_ADD_VALUE           int     = 0x4
	WORLD_SCALE           float32 = 15000
	MINIMAP_HUD_BUFF_SIZE int     = 0x80
)

var (
	HOOK    Hook.ProcessHook = Hook.HOOK
	MINIMAP Minimap
)

func Update() error {
	minimapObject, err := memory.ReadInt(HOOK.Process, HOOK.ModuleBaseAddr+offset.MINIMAPOBJECT)
	if err != nil {
		return err
	}

	if minimapObject <= 0 {
		return fmt.Errorf("error to find minimapObject")
	}

	minimapHUD, err := memory.ReadInt(HOOK.Process, minimapObject+offset.MINIMAPOBJECTHUD)
	if err != nil {
		return err
	}

	if minimapHUD <= 0 {
		return fmt.Errorf("error to find minimapHUD")
	}

	minimapHUDBuff, err := memory.Read(HOOK.Process, minimapHUD, MINIMAP_HUD_BUFF_SIZE)
	if err != nil {
		return err
	}

	MINIMAP.X = memory.Float32frombytes(minimapHUDBuff[offset.MINIMAPHUDPOS : offset.MINIMAPHUDPOS+4])
	MINIMAP.Y = memory.Float32frombytes(minimapHUDBuff[offset.MINIMAPHUDPOS+Y_ADD_VALUE : offset.MINIMAPHUDPOS+Y_ADD_VALUE+4])
	MINIMAP.Width = memory.Float32frombytes(minimapHUDBuff[offset.MINIMAPHUDSIZE : offset.MINIMAPHUDSIZE+4])
	MINIMAP.Height = memory.Float32frombytes(minimapHUDBuff[offset.MINIMAPHUDSIZE+Y_ADD_VALUE : offset.MINIMAPHUDSIZE+Y_ADD_VALUE+4])

	return nil
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
