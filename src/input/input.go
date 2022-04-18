package input

import (
	"fmt"
	"framework-memory-go/src/renderer"
	"framework-memory-go/src/win"
	"unsafe"
)

func IsKeyDown(key int) bool {
	return win.GetAsyncKeyState(int32(key)) < 0
}

func WasKeyPressed(key int) bool {
	state := int32(win.GetAsyncKeyState(int32(key))) & 0x8000
	if state != 0 {
		return true
	}
	return false
}

func MoveMouse(x, y int32) {
	win.SetCursorPos(x, y)
}

func PressRightClick() {
	mouseRightDown()
	mouseRightUp()
}

func PressLeftClick() {
	mouseLeftDown()
	mouseLefttUp()
}

func GetCursorPos() renderer.ScreenPosition {
	var pos renderer.ScreenPosition
	var point win.POINT
	if !win.GetCursorPos(&point) {
		fmt.Println("Error in GetCursorPos")
	}
	pos.X = float32(point.X)
	pos.Y = float32(point.Y)
	return pos
}

func mouseRightDown() {
	var input win.MOUSE_INPUT
	input.Type = win.INPUT_MOUSE
	input.Mi.DwFlags = win.MOUSEEVENTF_RIGHTDOWN | win.MOUSEEVENTF_VIRTUALDESK | win.MOUSEEVENTF_ABSOLUTE
	win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
}

func mouseRightUp() {
	var input win.MOUSE_INPUT
	input.Type = win.INPUT_MOUSE
	input.Mi.DwFlags = win.MOUSEEVENTF_RIGHTUP | win.MOUSEEVENTF_VIRTUALDESK | win.MOUSEEVENTF_ABSOLUTE
	win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
}

func mouseLeftDown() {
	var input win.MOUSE_INPUT
	input.Type = win.INPUT_MOUSE
	input.Mi.DwFlags = win.MOUSEEVENTF_LEFTDOWN
	win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
}

func mouseLefttUp() {
	var input win.MOUSE_INPUT
	input.Type = win.INPUT_MOUSE
	input.Mi.DwFlags = win.MOUSEEVENTF_LEFTUP
	win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
}
