package scripts

import (
	"framework-memory-go/src/unitmanager"
	"math"
	"strings"
)

const (
	HIGH_VALUE = 1000000
)

var Minionlist []string
var Towerlist []string
var Monsterlist []string

func init() {
	Minionlist = append(Minionlist, "apheliosturret")
	Minionlist = append(Minionlist, "elisespiderling")
	Minionlist = append(Minionlist, "ha_chaosminionmelee")
	Minionlist = append(Minionlist, "ha_chaosminionranged")
	Minionlist = append(Minionlist, "ha_chaosminionsiege")
	Minionlist = append(Minionlist, "ha_chaosminionsuper")
	Minionlist = append(Minionlist, "ha_orderminionmelee")
	Minionlist = append(Minionlist, "ha_orderminionranged")
	Minionlist = append(Minionlist, "ha_orderminionsiege")
	Minionlist = append(Minionlist, "ha_orderminionsuper")
	Minionlist = append(Minionlist, "kalistaspawn")
	Minionlist = append(Minionlist, "malzaharvoidling")
	Minionlist = append(Minionlist, "shacobox")
	Minionlist = append(Minionlist, "sru_chaosminionmelee")
	Minionlist = append(Minionlist, "sru_chaosminionranged")
	Minionlist = append(Minionlist, "sru_chaosminionsiege")
	Minionlist = append(Minionlist, "sru_chaosminionsuper")
	Minionlist = append(Minionlist, "sru_orderminionmelee")
	Minionlist = append(Minionlist, "sru_orderminionranged")
	Minionlist = append(Minionlist, "sru_orderminionsiege")
	Minionlist = append(Minionlist, "sru_orderminionsuper")
	Minionlist = append(Minionlist, "zyrathornplant")
	Minionlist = append(Minionlist, "zyragraspingplant")
	Minionlist = append(Minionlist, "zacrebirthbloblet")
	Minionlist = append(Minionlist, "yorickghoulmelee")
	Minionlist = append(Minionlist, "yorickbigghoul")
	Minionlist = append(Minionlist, "voidspawn")

	Towerlist = append(Towerlist, "azirsundisc")
	Towerlist = append(Towerlist, "bw_ap_chaosturret")
	Towerlist = append(Towerlist, "bw_ap_chaosturret2")
	Towerlist = append(Towerlist, "bw_ap_chaosturret3")
	Towerlist = append(Towerlist, "bw_ap_chaosturretrubble")
	Towerlist = append(Towerlist, "bw_ap_finn")
	Towerlist = append(Towerlist, "bw_ap_orderturret")
	Towerlist = append(Towerlist, "bw_ap_orderturret2")
	Towerlist = append(Towerlist, "bw_ap_orderturret3")
	Towerlist = append(Towerlist, "bw_ap_orderturretrubble")
	Towerlist = append(Towerlist, "ha_ap_chaosturret")
	Towerlist = append(Towerlist, "ha_ap_chaosturret2")
	Towerlist = append(Towerlist, "ha_ap_chaosturret3")
	Towerlist = append(Towerlist, "ha_ap_chaosturretrubble")
	Towerlist = append(Towerlist, "ha_ap_chaosturretshrine")
	Towerlist = append(Towerlist, "ha_ap_chaosturrettutorial")
	Towerlist = append(Towerlist, "ha_ap_ordershrineturret")
	Towerlist = append(Towerlist, "ha_ap_orderturret")
	Towerlist = append(Towerlist, "ha_ap_orderturret2")
	Towerlist = append(Towerlist, "ha_ap_orderturret3")
	Towerlist = append(Towerlist, "ha_ap_orderturretrubble")
	Towerlist = append(Towerlist, "ha_ap_orderturrettutorial")
	Towerlist = append(Towerlist, "preseason_turret_shield")
	Towerlist = append(Towerlist, "sruap_magecrystal")
	Towerlist = append(Towerlist, "sruap_orderinhibitor_rubble")
	Towerlist = append(Towerlist, "sruap_ordernexus")
	Towerlist = append(Towerlist, "sruap_ordernexus_rubble")
	Towerlist = append(Towerlist, "sruap_turret_chaos1")
	Towerlist = append(Towerlist, "sruap_turret_chaos2")
	Towerlist = append(Towerlist, "sruap_turret_chaos3")
	Towerlist = append(Towerlist, "sruap_turret_chaos4")
	Towerlist = append(Towerlist, "sruap_turret_chaos5")
	Towerlist = append(Towerlist, "sruap_turret_order1")
	Towerlist = append(Towerlist, "sruap_turret_order1_bot")
	Towerlist = append(Towerlist, "sruap_turret_order2")
	Towerlist = append(Towerlist, "sruap_turret_order3")
	Towerlist = append(Towerlist, "sruap_turret_order4")
	Towerlist = append(Towerlist, "sruap_turret_order5")
	Towerlist = append(Towerlist, "chaosinhibitor")
	Towerlist = append(Towerlist, "orderinhibitor")
	Towerlist = append(Towerlist, "sruap_orderinhibitor")
	Towerlist = append(Towerlist, "sruap_chaosinhibitor")
	Towerlist = append(Towerlist, "chaosnexus")
	Towerlist = append(Towerlist, "ordernexus")
	Towerlist = append(Towerlist, "sruap_chaosnexus")

	Monsterlist = append(Monsterlist, "assassinmode_objective_boss2")
	Monsterlist = append(Monsterlist, "doombotsbossteemo")
	Monsterlist = append(Monsterlist, "fiddlestickseffigy")
	Monsterlist = append(Monsterlist, "sru_baron")
	Monsterlist = append(Monsterlist, "kingporo")
	Monsterlist = append(Monsterlist, "sru_blue")
	Monsterlist = append(Monsterlist, "sru_bluemini")
	Monsterlist = append(Monsterlist, "sru_bluemini2")
	Monsterlist = append(Monsterlist, "sru_crab")
	Monsterlist = append(Monsterlist, "sru_dragon")
	Monsterlist = append(Monsterlist, "sru_dragon_air")
	Monsterlist = append(Monsterlist, "sru_dragon_chemtech")
	Monsterlist = append(Monsterlist, "sru_dragon_earth")
	Monsterlist = append(Monsterlist, "sru_dragon_elder")
	Monsterlist = append(Monsterlist, "sru_dragon_fire")
	Monsterlist = append(Monsterlist, "sru_dragon_hextech")
	Monsterlist = append(Monsterlist, "sru_dragon_ruined")
	Monsterlist = append(Monsterlist, "sru_dragon_water")
	Monsterlist = append(Monsterlist, "sru_gromp")
	Monsterlist = append(Monsterlist, "sru_krug")
	Monsterlist = append(Monsterlist, "sru_krugmini")
	Monsterlist = append(Monsterlist, "sru_krugminimini")
	Monsterlist = append(Monsterlist, "sru_murkwolf")
	Monsterlist = append(Monsterlist, "sru_murkwolfmini")
	Monsterlist = append(Monsterlist, "sru_razorbeak")
	Monsterlist = append(Monsterlist, "sru_razorbeakmini")
	Monsterlist = append(Monsterlist, "sru_red")
	Monsterlist = append(Monsterlist, "sru_redmini")
	Monsterlist = append(Monsterlist, "sru_riftherald")
	Monsterlist = append(Monsterlist, "sru_riftherald_mercenary")
}

func isMinion(name string) bool {
	for _, a := range Minionlist {
		if a == strings.ToLower(name) {
			return true
		}
	}
	return false
}

func isMonster(name string) bool {
	for _, a := range Monsterlist {
		if a == strings.ToLower(name) {
			return true
		}
	}
	return false
}

func isTower(name string) bool {
	for _, a := range Towerlist {
		if a == strings.ToLower(name) {
			return true
		}
	}
	return false
}

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

func GestBestTargetForUnits() (unitmanager.GameUnit, bool) {
	var lasthealth float32 = HIGH_VALUE
	gameunit := unitmanager.GameUnit{}
	// for _, element := range unitmanager.UNITMANAGER.Inhibitors {
	// 	if element.Team == unitmanager.LOCALPLAYER.Team {
	// 		continue
	// 	}
	// 	if !element.IsAlive {
	// 		continue
	// 	}
	// 	if !element.IsVisible {
	// 		continue
	// 	}
	// 	if !element.IsTargetable {
	// 		continue
	// 	}

	// 	if !inRange(element) {
	// 		continue
	// 	}
	// 	gameunit = element
	// 	break
	// }
	// if gameunit.Team != 0 {
	// 	return gameunit, true
	// }

	for _, element := range unitmanager.UNITMANAGER.Turrets {
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
		gameunit = element
		break
	}
	if gameunit.Team != 0 {
		return gameunit, true
	}
	for _, element := range unitmanager.UNITMANAGER.AllUnits {
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

	if gameunit.Team != 0 {
		return gameunit, true
	}
	return gameunit, false
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
