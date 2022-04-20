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
		log.Panic("not found process")
	}
	HOOK = processHook

	unitmanager.LoadUnitData()
	unitmanager.SpelltData()
	unitmanager.SummonerSpellData()
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

func (g *Game) Update() error {
	var wgg sync.WaitGroup
	if count == 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			count++
			wnproc := win.FindWindow(nil, hook.StringToUTF16PtrElseNil(NAME))
			r2, err := win.SetWindowLong(wnproc, win.GWL_EXSTYLE, win.WS_EX_COMPOSITED|win.WS_EX_LAYERED|win.WS_EX_TRANSPARENT|win.WS_EX_TOOLWINDOW|win.WS_EX_TOPMOST)
			if err != nil {
				fmt.Println("error in  setWindowLong: ", err, r2)
			}
			win.SetForegroundWindow(wnproc)
		}()

	}
	module.Update()
	wgg.Add(1)
	go func() {
		defer wgg.Done()
		scripts.UpdateOrbwalker()
	}()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var wgg sync.WaitGroup
	wgg.Add(1)
	go func() {
		defer wgg.Done()
		scripts.UpdateDrawings(screen)
	}()
	wgg.Wait()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}
