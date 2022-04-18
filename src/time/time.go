package time

import (
	Hook "framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
)

type Time struct {
	Second float32
}

var (
	hook Hook.ProcessHook
	TIME Time
)

func init() {
	hook = Hook.HOOK
}

func Update() error {
	value, err := memory.ReadFloat(hook.Process, int(hook.ModuleBaseAddr)+offset.GAMETIME)
	if err != nil {
		return err
	}
	TIME.Second = value
	return nil
}
