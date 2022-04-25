package scripts

import (
	"fmt"
	"framework-memory-go/src/gui"
	"framework-memory-go/src/minimap"
	"framework-memory-go/src/renderer"
	"framework-memory-go/src/time"
	"framework-memory-go/src/unitmanager"
	"image/color"
	"strings"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WARD_RANGE     = 900
	LAST_HIT_RANGE = 50
)

var meColor color.RGBA = color.RGBA{0, 162, 162, 1}
var enemyColor color.RGBA = color.RGBA{133, 162, 0, 1}
var wardColorRed color.RGBA = color.RGBA{255, 0, 0, 1}
var wardColorYellow color.RGBA = color.RGBA{231, 249, 0, 1}
var minionColorLastHit color.RGBA = color.RGBA{25, 0, 255, 1}

var wg sync.WaitGroup

func UpdateDrawings(screen *ebiten.Image) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		DrawChamps(screen)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		DrawUnits(screen)
	}()
	wg.Wait()
}

func DrawChamps(screen *ebiten.Image) {
	if !gui.CheckboxRanges.Value {
		return
	}
	for _, element := range unitmanager.UNITMANAGER.Champions {
		if element.Team != unitmanager.LOCALPLAYER.Team {
			if !element.IsAlive {
				continue
			}
			ran := element.AttackRange
			boundingradius := element.GameplayRadiusJson
			pos := element.Position
			if element.IsVisible {
				gui.DrawCircle(screen, pos, ran+boundingradius, 1.3, enemyColor)
				if element.Icon != nil {
					drawSpell(element, screen)
				}
				continue
			}
			missing := fmt.Sprintf("%.2f", time.TIME.Second-element.LastVisibleTime)
			rendererpos := renderer.WorldToScreen(renderer.RENDERER, pos.X, pos.Y, pos.Z)
			minimapos := minimap.MinimapToScreen(minimap.MINIMAP, pos.X, pos.Y, pos.Z)
			if element.Icon != nil {
				gui.DrawText(screen, int(rendererpos.X), int(rendererpos.Y), color.White, missing)
				gui.DrawImage(screen, float64(rendererpos.X), float64(rendererpos.Y), 0.5, 0.5, element.Icon, false)
				gui.DrawImage(screen, float64(minimapos.X)-10, float64(minimapos.Y)-10, 0.2, 0.2, element.Icon, false)
			}
			continue
		}
		if element.Name == unitmanager.LOCALPLAYER.Name {
			ran := element.AttackRange
			boundingradius := element.GameplayRadiusJson
			pos := element.Position
			gui.DrawCircle(screen, pos, ran+boundingradius, 4, meColor)
			continue
		}
	}
}

func DrawUnits(screen *ebiten.Image) {
	for _, element := range unitmanager.UNITMANAGER.AllUnits {
		if element.Team == unitmanager.LOCALPLAYER.Team {
			continue
		}

		if !element.IsAlive {
			continue
		}

		if isWard(strings.ToLower(element.Name)) {
			if !gui.CheckboxWards.Value {
				continue
			}
			rendererpos := renderer.WorldToScreen(renderer.RENDERER, element.Position.X, element.Position.Y, element.Position.Z)
			if element.Name == "JammerDevice" {
				gui.DrawText(screen, int(rendererpos.X), int(rendererpos.Y), color.White, element.Name)
				gui.DrawCircle(screen, element.Position, WARD_RANGE, 1.3, wardColorRed)
				gui.DrawCircleInMinimap(screen, element.Position, WARD_RANGE, 70, wardColorRed)
			} else {
				gui.DrawText(screen, int(rendererpos.X), int(rendererpos.Y), color.White, element.Name)
				gui.DrawCircle(screen, element.Position, WARD_RANGE, 1.3, wardColorYellow)
				gui.DrawCircleInMinimap(screen, element.Position, WARD_RANGE, 70, wardColorYellow)
			}
			continue
		}

		if isTrap(strings.ToLower(element.Name)) {
			if !gui.CheckboxTraps.Value {
				continue
			}
			rendererpos := renderer.WorldToScreen(renderer.RENDERER, element.Position.X, element.Position.Y, element.Position.Z)
			gui.DrawText(screen, int(rendererpos.X), int(rendererpos.Y), color.White, element.Name)
			gui.DrawCircle(screen, element.Position, element.GameplayRadiusJson, 1, wardColorRed)
			continue
		}

		if isMinion(strings.ToLower(element.Name)) {
			if !gui.CheckboxLastHit.Value {
				continue
			}
			armor := effectiveDamage(unitmanager.LOCALPLAYER.BaseAttack+unitmanager.LOCALPLAYER.BonusAttack, element.Armor+element.BonusArmor)
			if element.Health <= armor {
				gui.DrawCircle(screen, element.Position, LAST_HIT_RANGE, 2, minionColorLastHit)
			}
		}
	}
}

var iconSize float32 = 28
var yOffset float32 = iconSize * 2

func drawSpell(gameUnit unitmanager.GameUnit, screen *ebiten.Image) {
	if !gui.CheckboxSpells.Value {
		return
	}
	rendererpos := renderer.WorldToScreen(renderer.RENDERER, gameUnit.Position.X, gameUnit.Position.Y, gameUnit.Position.Z)
	drawY := rendererpos.Y + yOffset - 150
	xOffset := -yOffset - 20
	for _, element := range gameUnit.Spells {
		levelled := element.Level >= 1
		remaining := element.ReadyAtSeconds - time.TIME.Second
		ready := remaining <= 0
		if element.Icon == nil {
			continue
		}
		if !levelled || !ready {
			gui.DrawImage(screen, float64(rendererpos.X)+float64(xOffset), float64(drawY), 0.40, 0.40, element.Icon, true)
		} else {
			gui.DrawImage(screen, float64(rendererpos.X)+float64(xOffset), float64(drawY), 0.40, 0.40, element.Icon, false)
		}
		if levelled && !ready {
			gui.DrawText(screen, int(rendererpos.X)+int(xOffset), int(drawY), color.White, fmt.Sprintf("%.1f", remaining))
		}
		xOffset += iconSize
	}
}
