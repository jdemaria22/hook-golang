package time

import (
	"framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
)

type Time struct {
	second float32
}

func Update(hook hook.ProcessHook) (Time, error) {
	var time Time
	value, err := memory.ReadFloat(hook, int(hook.ModuleBaseAddr)+offset.GAMETIME)
	if err != nil {
		return time, err
	}
	time.second = value
	return time, nil
}
