package main

import (
	"fmt"
	"framework-memory-go/src/hook"
	"framework-memory-go/src/module"
	"framework-memory-go/src/scripts"
	"framework-memory-go/src/unitmanager"
	"framework-memory-go/src/win"
	"log"
	"math/rand"
	"strconv"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	HOOK  hook.ProcessHook
	count = 0
	NAME  string
)

const (
	screenWidth  = 1920
	screenHeight = 1080
	RANDOM       = 10000
	MAX_TPS      = 1000
)

func init() {
	processHook, err := hook.GetHook()
	if err != nil {
		fmt.Println(err)
	}
	HOOK = processHook

	unitmanager.LoadUnitData()
	unitmanager.SpelltData()
}

func main() {
	NAME = strconv.Itoa(rand.Intn(RANDOM))
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(NAME)
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowDecorated(false)
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetInitFocused(true)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetMaxTPS(60)

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
}

func NewGame() *Game {
	return nil
}

/*
	Update de la memoria.
	entre 0 y 2 miliseconds dura la ejecución
*/
func (g *Game) Update() error {
	if count == 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			count++
			yogur := win.FindWindow(nil, hook.StringToUTF16PtrElseNil(NAME))
			style := win.GetWindowLong(yogur, win.GWL_EXSTYLE)
			r2, err := win.SetWindowLong(yogur, win.GWL_EXSTYLE, style&^(win.WS_EX_DLGMODALFRAME|win.WS_EX_WINDOWEDGE|win.WS_EX_CLIENTEDGE|win.WS_EX_STATICEDGE))
			if err != nil {
				fmt.Println("error in  setWindowLong: ", err, r2)
			}
			r2, err = win.SetWindowLong(yogur, win.GWL_EXSTYLE, win.WS_EX_OVERLAPPEDWINDOW|win.WS_EX_LAYERED|win.WS_EX_TRANSPARENT|win.WS_EX_TOPMOST)
			win.SetForegroundWindow(yogur)
			if err != nil {
				fmt.Println("error in  setWindowLong: ", err, r2)
			}
		}()

	}
	module.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scripts.UpdateDrawings(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
