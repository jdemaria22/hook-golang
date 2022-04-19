package scripts

import (
	"framework-memory-go/src/unitmanager"
	"math"
	"strings"
)

const (
	HIGH_VALUE = 1000000
)

var minionlist []string
var towerlist []string
var monsterlist []string

func init() {
	minionlist = append(minionlist, "apheliosturret")
	minionlist = append(minionlist, "elisespiderling")
	minionlist = append(minionlist, "ha_chaosminionmelee")
	minionlist = append(minionlist, "ha_chaosminionranged")
	minionlist = append(minionlist, "ha_chaosminionsiege")
	minionlist = append(minionlist, "ha_chaosminionsuper")
	minionlist = append(minionlist, "ha_orderminionmelee")
	minionlist = append(minionlist, "ha_orderminionranged")
	minionlist = append(minionlist, "ha_orderminionsiege")
	minionlist = append(minionlist, "ha_orderminionsuper")
	minionlist = append(minionlist, "kalistaspawn")
	minionlist = append(minionlist, "malzaharvoidling")
	minionlist = append(minionlist, "shacobox")
	minionlist = append(minionlist, "sru_chaosminionmelee")
	minionlist = append(minionlist, "sru_chaosminionranged")
	minionlist = append(minionlist, "sru_chaosminionsiege")
	minionlist = append(minionlist, "sru_chaosminionsuper")
	minionlist = append(minionlist, "sru_orderminionmelee")
	minionlist = append(minionlist, "sru_orderminionranged")
	minionlist = append(minionlist, "sru_orderminionsiege")
	minionlist = append(minionlist, "sru_orderminionsuper")
	minionlist = append(minionlist, "zyrathornplant")
	minionlist = append(minionlist, "zyragraspingplant")
	minionlist = append(minionlist, "zacrebirthbloblet")
	minionlist = append(minionlist, "yorickghoulmelee")
	minionlist = append(minionlist, "yorickbigghoul")
	minionlist = append(minionlist, "voidspawn")

	towerlist = append(towerlist, "azirsundisc")
	towerlist = append(towerlist, "bw_ap_chaosturret")
	towerlist = append(towerlist, "bw_ap_chaosturret2")
	towerlist = append(towerlist, "bw_ap_chaosturret3")
	towerlist = append(towerlist, "bw_ap_chaosturretrubble")
	towerlist = append(towerlist, "bw_ap_finn")
	towerlist = append(towerlist, "bw_ap_orderturret")
	towerlist = append(towerlist, "bw_ap_orderturret2")
	towerlist = append(towerlist, "bw_ap_orderturret3")
	towerlist = append(towerlist, "bw_ap_orderturretrubble")
	towerlist = append(towerlist, "ha_ap_chaosturret")
	towerlist = append(towerlist, "ha_ap_chaosturret2")
	towerlist = append(towerlist, "ha_ap_chaosturret3")
	towerlist = append(towerlist, "ha_ap_chaosturretrubble")
	towerlist = append(towerlist, "ha_ap_chaosturretshrine")
	towerlist = append(towerlist, "ha_ap_chaosturrettutorial")
	towerlist = append(towerlist, "ha_ap_ordershrineturret")
	towerlist = append(towerlist, "ha_ap_orderturret")
	towerlist = append(towerlist, "ha_ap_orderturret2")
	towerlist = append(towerlist, "ha_ap_orderturret3")
	towerlist = append(towerlist, "ha_ap_orderturretrubble")
	towerlist = append(towerlist, "ha_ap_orderturrettutorial")
	towerlist = append(towerlist, "preseason_turret_shield")
	towerlist = append(towerlist, "sruap_magecrystal")
	towerlist = append(towerlist, "sruap_orderinhibitor_rubble")
	towerlist = append(towerlist, "sruap_ordernexus")
	towerlist = append(towerlist, "sruap_ordernexus_rubble")
	towerlist = append(towerlist, "sruap_turret_chaos1")
	towerlist = append(towerlist, "sruap_turret_chaos2")
	towerlist = append(towerlist, "sruap_turret_chaos3")
	towerlist = append(towerlist, "sruap_turret_chaos4")
	towerlist = append(towerlist, "sruap_turret_chaos5")
	towerlist = append(towerlist, "sruap_turret_order1")
	towerlist = append(towerlist, "sruap_turret_order1_bot")
	towerlist = append(towerlist, "sruap_turret_order2")
	towerlist = append(towerlist, "sruap_turret_order3")
	towerlist = append(towerlist, "sruap_turret_order4")
	towerlist = append(towerlist, "sruap_turret_order5")
	towerlist = append(towerlist, "chaosinhibitor")
	towerlist = append(towerlist, "orderinhibitor")
	towerlist = append(towerlist, "sruap_orderinhibitor")
	towerlist = append(towerlist, "sruap_chaosinhibitor")
	towerlist = append(towerlist, "chaosnexus")
	towerlist = append(towerlist, "ordernexus")
	towerlist = append(towerlist, "sruap_chaosnexus")

	monsterlist = append(monsterlist, "assassinmode_objective_boss2")
	monsterlist = append(monsterlist, "doombotsbossteemo")
	monsterlist = append(monsterlist, "fiddlestickseffigy")
	monsterlist = append(monsterlist, "sru_baron")
	monsterlist = append(monsterlist, "kingporo")
	monsterlist = append(monsterlist, "sru_blue")
	monsterlist = append(monsterlist, "sru_bluemini")
	monsterlist = append(monsterlist, "sru_bluemini2")
	monsterlist = append(monsterlist, "sru_crab")
	monsterlist = append(monsterlist, "sru_dragon")
	monsterlist = append(monsterlist, "sru_dragon_air")
	monsterlist = append(monsterlist, "sru_dragon_chemtech")
	monsterlist = append(monsterlist, "sru_dragon_earth")
	monsterlist = append(monsterlist, "sru_dragon_elder")
	monsterlist = append(monsterlist, "sru_dragon_fire")
	monsterlist = append(monsterlist, "sru_dragon_hextech")
	monsterlist = append(monsterlist, "sru_dragon_ruined")
	monsterlist = append(monsterlist, "sru_dragon_water")
	monsterlist = append(monsterlist, "sru_gromp")
	monsterlist = append(monsterlist, "sru_krug")
	monsterlist = append(monsterlist, "sru_krugmini")
	monsterlist = append(monsterlist, "sru_krugminimini")
	monsterlist = append(monsterlist, "sru_murkwolf")
	monsterlist = append(monsterlist, "sru_murkwolfmini")
	monsterlist = append(monsterlist, "sru_razorbeak")
	monsterlist = append(monsterlist, "sru_razorbeakmini")
	monsterlist = append(monsterlist, "sru_red")
	monsterlist = append(monsterlist, "sru_redmini")
	monsterlist = append(monsterlist, "sru_riftherald")
	monsterlist = append(monsterlist, "sru_riftherald_mercenary")
}

func isMinion(name string) bool {
	for _, a := range minionlist {
		if a == strings.ToLower(name) {
			return true
		}
	}
	return false
}

func isMonster(name string) bool {
	for _, a := range monsterlist {
		if a == strings.ToLower(name) {
			return true
		}
	}
	return false
}

func isTower(name string) bool {
	for _, a := range towerlist {
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
	for _, element := range unitmanager.UNITMANAGER.Inhibitors {
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
	for _, element := range unitmanager.UNITMANAGER.Minions {
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
