package unitmanager

import (
	"fmt"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
)

type AiManager struct {
	IsMoving           bool
	IsDashing          bool
	ServerPos          GamePosition
	NavBegin           GamePosition
	NavEnd             GamePosition
	OwnPosition        GamePosition
	ClickRightPosition GamePosition
	CastPosition       GamePosition
	Velocity           GamePosition
	NavigationPath     []GamePosition
	MoveSpeed          float32
	DashSpeed          float32
}

var AI_MANAGER_BASE = 0x4
var AI_VALUE_TO_MOVE = 0x8

func UpdateAimanager(data []byte) AiManager {
	offsetSubBase := offset.AIMANAGER - AI_MANAGER_BASE

	off := offsetSubBase + AI_VALUE_TO_MOVE
	offadd := data[offsetSubBase+AI_MANAGER_BASE]
	objectAiManagerMoveTwoBytes := memory.Int32frombytes(data[off+int(offadd)*AI_MANAGER_BASE:])

	objectAiManagerSubBase := memory.Int32frombytes(data[offsetSubBase:])

	objectAiManagerMoveTwoBytes ^= ^objectAiManagerSubBase

	base, err := memory.ReadInt(HOOK.Process, int(objectAiManagerMoveTwoBytes)+AI_VALUE_TO_MOVE)
	if err != nil {
		fmt.Println("Error in UpdateAimanager.info ", err)
		return AiManager{}
	}
	baseBuff, err := memory.Read(HOOK.Process, base, 0x1100)
	aiManager := AiManager{}
	aiManager.IsMoving = baseBuff[offset.AIMANAGERISMOVING] != 0
	aiManager.IsDashing = baseBuff[offset.AIMANAGERISDASHING] != 0

	aiManager.ServerPos.X = memory.Float32frombytes(baseBuff[offset.AIMANAGERSERVPOSITION:])
	aiManager.ServerPos.Y = memory.Float32frombytes(baseBuff[offset.AIMANAGERSERVPOSITION+0x4:])
	aiManager.ServerPos.Z = memory.Float32frombytes(baseBuff[offset.AIMANAGERSERVPOSITION+0x8:])

	aiManager.NavBegin.X = memory.Float32frombytes(baseBuff[offset.AIMANAGERSTARTPATH:])
	aiManager.NavBegin.Y = memory.Float32frombytes(baseBuff[offset.AIMANAGERSTARTPATH+0x4:])
	aiManager.NavBegin.Z = memory.Float32frombytes(baseBuff[offset.AIMANAGERSTARTPATH+0x8:])

	aiManager.NavEnd.X = memory.Float32frombytes(baseBuff[offset.AIMANAGERENDPATH:])
	aiManager.NavEnd.Y = memory.Float32frombytes(baseBuff[offset.AIMANAGERENDPATH+0x4:])
	aiManager.NavEnd.Z = memory.Float32frombytes(baseBuff[offset.AIMANAGERENDPATH+0x8:])

	aiManager.OwnPosition.X = memory.Float32frombytes(baseBuff[offset.AIMANAGEROWNPOSITION:])
	aiManager.OwnPosition.Y = memory.Float32frombytes(baseBuff[offset.AIMANAGEROWNPOSITION+0x4:])
	aiManager.OwnPosition.Z = memory.Float32frombytes(baseBuff[offset.AIMANAGEROWNPOSITION+0x8:])

	aiManager.ClickRightPosition.X = memory.Float32frombytes(baseBuff[offset.AIMANAGERCLICKRIGHTPOSITION:])
	aiManager.ClickRightPosition.Y = memory.Float32frombytes(baseBuff[offset.AIMANAGERCLICKRIGHTPOSITION+0x4:])
	aiManager.ClickRightPosition.Z = memory.Float32frombytes(baseBuff[offset.AIMANAGERCLICKRIGHTPOSITION+0x8:])

	aiManager.CastPosition.X = memory.Float32frombytes(baseBuff[offset.AIMANAGERCASTPOSITION:])
	aiManager.CastPosition.Y = memory.Float32frombytes(baseBuff[offset.AIMANAGERCASTPOSITION+0x4:])
	aiManager.CastPosition.Z = memory.Float32frombytes(baseBuff[offset.AIMANAGERCASTPOSITION+0x8:])

	aiManager.Velocity.X = memory.Float32frombytes(baseBuff[offset.AIMANAGERVELOCITY:])
	aiManager.Velocity.Y = memory.Float32frombytes(baseBuff[offset.AIMANAGERVELOCITY+0x4:])
	aiManager.Velocity.Z = memory.Float32frombytes(baseBuff[offset.AIMANAGERVELOCITY+0x8:])

	aiManager.MoveSpeed = memory.Float32frombytes(baseBuff[offset.AIMANAGERMOVESPEED:])

	aiManager.DashSpeed = memory.Float32frombytes(baseBuff[offset.AIMANAGERDASHSPEED:])
	aiManager.NavigationPath = navigationPath(baseBuff)
	return aiManager
}

var AiManagerPointerPathStart = 0x1E4
var AiManagerPointerPathEnd = 0x1E8

func navigationPath(data []byte) []GamePosition {
	aiManagerPointerPathStart := memory.Int32frombytes(data[AiManagerPointerPathStart:])
	if aiManagerPointerPathStart == 0 {
		return []GamePosition{}
	}

	aiManagerPointerPathEnd := memory.Int32frombytes(data[AiManagerPointerPathEnd:])

	if aiManagerPointerPathEnd == 0 {
		return []GamePosition{}
	}
	numsegments := aiManagerPointerPathEnd - aiManagerPointerPathStart
	if numsegments == 0 {
		return []GamePosition{}
	}
	dataBuff, err := memory.Read(HOOK.Process, int(aiManagerPointerPathStart), int(numsegments))
	if err != nil {
		fmt.Println(err)
		return []GamePosition{}
	}

	navigationPath := []GamePosition{}
	gamePosition := GamePosition{}
	for i := 0; i < int(numsegments); i += 0xC {
		gamePosition.X = memory.Float32frombytes(dataBuff[i:])
		gamePosition.Y = memory.Float32frombytes(dataBuff[i+0x4:])
		gamePosition.Z = memory.Float32frombytes(dataBuff[i+0x8:])
		navigationPath = append(navigationPath, gamePosition)
	}

	return navigationPath
}
