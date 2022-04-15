package scripts

import (
	"framework-memory-go/src/gui"
	"framework-memory-go/src/unitmanager"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateDrawings(screen *ebiten.Image) {
	DrawChamps(screen)
	DrawMinions(screen)
}

func DrawChamps(screen *ebiten.Image) {

	purpleClr := color.RGBA{255, 0, 255, 255}
	for _, element := range unitmanager.UNITMANAGER.Champions {
		ran := element.AttackRange
		pos := element.Position
		gui.DrawCircle(screen, pos, ran+65, 4, purpleClr)
	}
}

func DrawMinions(screen *ebiten.Image) {
	purpleClr := color.RGBA{255, 0, 255, 255}
	for _, element := range unitmanager.UNITMANAGER.Minions {
		ran := element.AttackRange
		pos := element.Position
		gui.DrawCircle(screen, pos, ran+65, 4, purpleClr)
	}
}
