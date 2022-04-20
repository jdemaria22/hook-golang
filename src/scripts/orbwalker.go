package scripts

import (
	"framework-memory-go/src/input"
	"framework-memory-go/src/renderer"
	t "framework-memory-go/src/time"
	"framework-memory-go/src/unitmanager"
	"framework-memory-go/src/utils"
	"framework-memory-go/src/win"
	"strings"

	"math/rand"
	"time"
)

const (
	LETHAL_TEMPO = "ASSETS/Perks/Styles/Precision/LethalTempo/LethalTempo.lua"
	VK_KEY       = 0x56
)

var isOwrbwalking = false
var isLaneClearing = false

func UpdateOrbwalker() {
	if utils.UTILS.IsChatOpen {
		return
	}
	if input.IsKeyDown(win.VK_SPACE) {
		target, ok := GestBestTarget()
		if ok && canAttack() && !isOwrbwalking {
			isOwrbwalking = true
			oldpost := input.GetCursorPos()
			targetScreenPos := renderer.WorldToScreen(renderer.RENDERER, target.Position.X, target.Position.Y, target.Position.Z)
			input.MoveMouse(int32(targetScreenPos.X), int32(targetScreenPos.Y))
			input.PressRightClick()
			attackTimer = time.Now().UnixMilli()
			lastMoveCommandT = time.Now().UnixMilli() + int64(getWindUpTime())
			time.Sleep(25 * time.Millisecond)
			input.MoveMouse(int32(oldpost.X), int32(oldpost.Y))
			isOwrbwalking = false
			return
		}
		if canMove() {
			input.PressRightClick()
			lastMoveCommandT = randomNumber() + time.Now().UnixMilli()
			return
		}
	}
	if input.IsKeyDown(VK_KEY) {
		target, ok := GestBestTargetForUnits()
		if ok && canAttack() && !isLaneClearing {
			isLaneClearing = true
			oldpost := input.GetCursorPos()
			targetScreenPos := renderer.WorldToScreen(renderer.RENDERER, target.Position.X, target.Position.Y, target.Position.Z)
			input.MoveMouse(int32(targetScreenPos.X), int32(targetScreenPos.Y))
			input.PressRightClick()
			attackTimer = time.Now().UnixMilli()
			lastMoveCommandT = time.Now().UnixMilli() + int64(getWindUpTime())
			time.Sleep(25 * time.Millisecond)
			input.MoveMouse(int32(oldpost.X), int32(oldpost.Y))
			isLaneClearing = false
			return
		}
		if canMove() {
			input.PressRightClick()
			lastMoveCommandT = randomNumber() + time.Now().UnixMilli()
			return
		}
	}
}

var attackTimer int64 = 0
var lastMoveCommandT int64 = 0
var ping int64 = 45 / 2
var min int64 = 150
var max int64 = 300

func canAttack() bool {
	return attackTimer+int64(getAttackDelay())+ping < time.Now().UnixMilli()
}

func getAttackDelay() float32 {
	return 1000 / getAttackSpeed()
}

func getAttackSpeed() float32 {
	attackSpeed := unitmanager.LOCALPLAYER.AttackSpeed * unitmanager.LOCALPLAYER.AttackSpeedMulti
	if attackSpeed <= 2.5 {
		return attackSpeed
	}
	if isLethalTempoActive(unitmanager.LOCALPLAYER.Buffs) {
		return attackSpeed
	}
	return 2.5
}

func getWindUpTime() float32 {
	return (1 / getAttackSpeed() * 1000) * unitmanager.LOCALPLAYER.BasicAtkWindup
}

func canMove() bool {
	return lastMoveCommandT < time.Now().UnixMilli()
}

func randomNumber() int64 {
	return rand.Int63n(max-min) + min
}

func isLethalTempoActive(buffs []unitmanager.Buff) bool {
	for _, element := range buffs {
		if !strings.Contains(element.Name, LETHAL_TEMPO) {
			continue
		}
		if !(element.EndTime > t.TIME.Second) {
			continue
		}
		if element.Count != 6 {
			continue
		}
		return true
	}
	return false
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}
	return false
}
