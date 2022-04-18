package unitmanager

import (
	"fmt"
	"framework-memory-go/src/memory"
	"framework-memory-go/src/offset"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Spell struct {
	Name               string
	ReadyAtSeconds     float32
	Level              int32
	Value              float32
	Icon               *ebiten.Image
	IconName           string
	Flags              int
	Delay              float32
	CastRange          float32
	CastRadius         float32
	Width              float32
	Height             float32
	Speed              float32
	TravelTime         float32
	ProjectDestination bool
	Type               string
}

const (
	DATA_SIZE            = 0x150
	SIZE_NAME_SPELL      = 50
	SpellSlotLevel       = 0x1C
	SpellSlotTime        = 0x24
	SpellSlotDamage      = 0x94
	SpellSlotSpellInfo   = 0x120
	SpellInfoSpellData   = 0x44
	SpellDataSpellName   = 0x6C
	SpellDataMissileName = 0x78
)

func UpdateSpell(gameUnit GameUnit, data []byte, deep bool) [6]Spell {
	for i := 0; i < len(gameUnit.Spells); i++ {
		address := memory.Int32frombytes(data[offset.OBJSPELLBOOK+(i*4):])
		if address == 0 {
			fmt.Println("Error in address UpdateSpell")
			continue
		}
		databuff, err := memory.Read(HOOK.Process, int(address), DATA_SIZE)
		if err != nil {
			fmt.Println("Error in databuff ", err)
			continue
		}
		gameUnit.Spells[i].ReadyAtSeconds = memory.Float32frombytes(databuff[SpellSlotTime:])
		gameUnit.Spells[i].Level = memory.Int32frombytes(databuff[SpellSlotLevel:])
		gameUnit.Spells[i].Value = memory.Float32frombytes(databuff[SpellSlotDamage:])
		if deep {
			gameUnit.Spells[i] = deepLoad(databuff, gameUnit.Spells[i])
		}
	}
	return gameUnit.Spells
}

func deepLoad(data []byte, spell Spell) Spell {
	infoPointer := memory.Int32frombytes(data[SpellSlotSpellInfo:])
	if infoPointer == 0 {
		fmt.Println("Error in Infopointer")
	}
	dataPointer, _ := memory.ReadInt(HOOK.Process, int(infoPointer+SpellInfoSpellData))
	if dataPointer == 0 {
		fmt.Println("Error in dataPointer")
	}
	namePointer, _ := memory.ReadInt(HOOK.Process, int(dataPointer+SpellDataSpellName))
	if namePointer == 0 {
		fmt.Println("Error in namePointer")
	}
	val, err := memory.Read(HOOK.Process, int(namePointer), SIZE_NAME_SPELL)
	if err != nil {
		fmt.Println("Error in namePointer merca")
	}
	spell.Name = memory.CopyString(val)
	// fmt.Println(spell.Name)
	spell = addSpellInfoFromJson(spell)
	spell = loadIconSpell(spell)
	return spell
}

func addSpellInfoFromJson(spell Spell) Spell {
	for i := 0; i < len(SPELL_DATA); i++ {
		if strings.ToLower(SPELL_DATA[i].Name) == strings.ToLower(spell.Name) {
			spell.IconName = SPELL_DATA[i].Icon
			spell.Flags = SPELL_DATA[i].Flags
			spell.Delay = SPELL_DATA[i].Delay
			spell.CastRange = SPELL_DATA[i].CastRange
			spell.CastRadius = SPELL_DATA[i].CastRadius
			spell.Width = SPELL_DATA[i].Width
			spell.Height = SPELL_DATA[i].Height
			spell.Speed = SPELL_DATA[i].Speed
			spell.TravelTime = SPELL_DATA[i].TravelTime
			spell.ProjectDestination = SPELL_DATA[i].ProjectDestination
			spell.Type = SPELL_DATA[i].Type
			return spell
		}
	}
	for i := 0; i < len(SUMMONER_SPELL_DATA); i++ {
		if strings.ToLower(SUMMONER_SPELL_DATA[i].Name) == strings.ToLower(spell.Name) {
			spell.IconName = SUMMONER_SPELL_DATA[i].Icon
			spell.Flags = SUMMONER_SPELL_DATA[i].Flags
			spell.Delay = SUMMONER_SPELL_DATA[i].Delay
			spell.CastRange = SUMMONER_SPELL_DATA[i].CastRange
			spell.CastRadius = SUMMONER_SPELL_DATA[i].CastRadius
			spell.Width = SUMMONER_SPELL_DATA[i].Width
			spell.Height = SUMMONER_SPELL_DATA[i].Height
			spell.Speed = SUMMONER_SPELL_DATA[i].Speed
			spell.TravelTime = SUMMONER_SPELL_DATA[i].TravelTime
			spell.ProjectDestination = SUMMONER_SPELL_DATA[i].ProjectDestination
			spell.Type = SUMMONER_SPELL_DATA[i].Type
			return spell
		}
	}
	return spell
}

func loadIconSpell(spell Spell) Spell {
	img, _, err := ebitenutil.NewImageFromFile(resourcepath + "/icons_spells/" + strings.ToLower(spell.IconName) + ".png")
	if err != nil {
		fmt.Println("err in load icon spell : ", err)
		return spell
	}
	spell.Icon = img
	return spell
}
