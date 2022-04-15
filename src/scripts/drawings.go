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
	// for i := 0; i < len(unitmanager.UNITMANAGER.Champions); i++ {
	// 	ran := unitmanager.UNITMANAGER.Champions[i].AttackRange
	// 	pos := unitmanager.UNITMANAGER.Champions[i].Position
	// 	gui.DrawCircle(screen, pos, ran+65, 5, purpleClr)
	// }
	for _, element := range unitmanager.UNITMANAGER.Champions {
		ran := element.AttackRange
		pos := element.Position
		gui.DrawCircle(screen, pos, ran+65, 4, purpleClr)
	}
}

func DrawMinions(screen *ebiten.Image) {
	purpleClr := color.RGBA{255, 0, 255, 255}
	for i := 0; i < len(unitmanager.UNITMANAGER.Minions); i++ {
		ran := unitmanager.UNITMANAGER.Minions[i].AttackRange
		pos := unitmanager.UNITMANAGER.Minions[i].Position
		gui.DrawCircle(screen, pos, ran+65, 4, purpleClr)
	}
}
