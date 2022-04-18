package scripts

import (
	"framework-memory-go/src/unitmanager"
	"math"
)

const (
	HIGH_VALUE = 1000000
)

func GestBestTarget() (unitmanager.GameUnit, bool) {
	var lasthealth float32 = HIGH_VALUE
	gameunit := unitmanager.GameUnit{}
	for _, element := range unitmanager.UNITMANAGER.Champions {
		if element.Team == unitmanager.LOCALPLAYER.Team {
			continue
		}
		if !element.IsAlive {
			continue
		}
		if !element.IsVisible {
			continue
		}
		if !element.IsTargetable {
			continue
		}
		if !inRange(element) {
			continue
		}
		if lasthealth >= element.Health {
			lasthealth = element.Health
			gameunit = element
		}
	}
	if gameunit.Team == 0 {
		return gameunit, false
	}
	return gameunit, true
}

func inRange(gameunit unitmanager.GameUnit) bool {
	entityradius := gameunit.GameplayRadiusJson * gameunit.SizeMultiplier
	championradius := gameunit.GameplayRadiusJson * unitmanager.LOCALPLAYER.SizeMultiplier
	return distanceBetweenTargets3D(unitmanager.LOCALPLAYER.Position, gameunit.Position)-float64(entityradius) < float64(unitmanager.LOCALPLAYER.AttackRange)+float64(championradius)
}

func distanceBetweenTargets3D(position1 unitmanager.GamePosition, position2 unitmanager.GamePosition) float64 {
	pow := math.Pow(float64(position1.X)-float64(position2.X), 2) + math.Pow(float64(position1.Y)-float64(position2.Y), 2) + math.Pow(float64(position1.Z)-float64(position2.Z), 2)
	return math.Abs(math.Pow(pow, 0.5))
}
