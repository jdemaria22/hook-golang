package hook

import (
	"fmt"
	"framework-memory-go/src/win"
	"syscall"
)

type ProcessHook struct {
	Window         win.HWND
	Pid            uint32
	Process        win.HANDLE
	ModuleBaseAddr int
}

var (
	HOOK ProcessHook
)

func init() {
	hook, err := GetHook()
	if err != nil {
		fmt.Println("Error in init unitmanager", err)
		return
	}
	HOOK = hook
}

const WINDOW_NAME = "RiotWindowClass"
const BASE_MODULE_NAME = "League of Legends.exe"
const PROCESS_ALL_ACCES uint32 = (0x000F0000 | 0x00100000 | 0xFFF)

func GetHook() (ProcessHook, error) {
	var processHook ProcessHook
	window := win.FindWindow(StringToUTF16PtrElseNil("RiotWindowClass"), nil)
	processHook.Window = window

	pid := win.GetWindowPidThreadProcessId(processHook.Window)
	processHook.Pid = pid

	process, err := win.OpenProcess(PROCESS_ALL_ACCES, false, processHook.Pid)
	if err != nil {
		return processHook, err
	}
	processHook.Process = process

	moduleBaseAddr, err := win.EnumProcessModule(process)
	if err != nil {
		return processHook, err
	}
	processHook.ModuleBaseAddr = int(moduleBaseAddr)
	return processHook, nil
}

func StringToUTF16PtrElseNil(str string) *uint16 {
	if str == "" {
		return nil
	}
	return syscall.StringToUTF16Ptr(str)
}
