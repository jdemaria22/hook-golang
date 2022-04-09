package main

import (
	"fmt"
	"framework-memory-go/src/gui"
	"framework-memory-go/src/hook"
	"framework-memory-go/src/module"
	"framework-memory-go/src/win"
	"image/color"
	"log"
	"math"
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
}

func main() {
	NAME = strconv.Itoa(rand.Intn(RANDOM))
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(NAME)
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowDecorated(false)
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetMaxTPS(60)
	ebiten.SetVsyncEnabled(false)
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
	if count == 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			count++
			yogur := win.FindWindow(nil, hook.StringToUTF16PtrElseNil(NAME))
			r2, err := win.SetWindowLong(yogur, win.GWL_EXSTYLE, win.WS_EX_COMPOSITED|win.WS_EX_LAYERED|win.WS_EX_TRANSPARENT|win.WS_EX_TOOLWINDOW|win.WS_EX_TOPMOST)
			if err != nil {
				fmt.Println("error in  setWindowLong: ", err, r2)
			}
		}()
	}

	return nil
}

func (g *Game) drawCircle(screen *ebiten.Image, x, y, radius int, clr color.Color) {
	radius64 := float64(radius)
	minAngle := math.Acos(1 - 1/radius64)

	for angle := float64(0); angle <= 360; angle += minAngle {
		xDelta := radius64 * math.Cos(angle)
		yDelta := radius64 * math.Sin(angle)

		x1 := int(math.Round(float64(x) + xDelta))
		y1 := int(math.Round(float64(y) + yDelta))

		screen.Set(x1, y1, clr)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	module.Update()
	screen.Clear()
	gui.DrawChamps(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
