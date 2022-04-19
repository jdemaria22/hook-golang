package utils

import (
	"fmt"
	Hook "framework-memory-go/src/hook"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
)

type Utils struct {
	IsChatOpen bool
}

var (
	HOOK  Hook.ProcessHook = Hook.HOOK
	UTILS Utils
)

func Update() error {
	chatInstance, err := memory.ReadInt(HOOK.Process, HOOK.ModuleBaseAddr+offset.CHAT)
	if err != nil {
		fmt.Println("Error in AIMinionClient ", err)
	}

	chatOpenVal, err := memory.Read(HOOK.Process, chatInstance+offset.CHATISOPEN, 1)
	UTILS.IsChatOpen = chatOpenVal[0] != 0
	return nil
}
