package unitmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type UnitData []struct {
	Name                 string   `json:"name"`
	HealthBarHeight      float32  `json:"healthBarHeight"`
	BaseMoveSpeed        float32  `json:"baseMoveSpeed"`
	AttackRange          float32  `json:"attackRange"`
	AttackSpeed          float32  `json:"attackSpeed"`
	AttackSpeedRatio     float32  `json:"attackSpeedRatio"`
	AcquisitionRange     float32  `json:"acquisitionRange"`
	SelectionRadius      float32  `json:"selectionRadius"`
	PathingRadius        float32  `json:"pathingRadius"`
	GameplayRadius       float32  `json:"gameplayRadius"`
	BasicAtkMissileSpeed float32  `json:"basicAtkMissileSpeed"`
	BasicAtkWindup       float32  `json:"basicAtkWindup"`
	Tags                 []string `json:"tags"`
}

type SpellData []struct {
	Name               string  `json:"name"`
	Icon               string  `json:"icon"`
	Flags              int     `json:"flags"`
	Delay              float32 `json:"delay"`
	CastRange          float32 `json:"castRange"`
	CastRadius         float32 `json:"castRadius"`
	Width              float32 `json:"width"`
	Height             float32 `json:"height"`
	Speed              float32 `json:"speed"`
	TravelTime         float32 `json:"travelTime"`
	ProjectDestination bool    `json:"projectDestination"`
	Type               string  `json:"Type"`
}

var UNIT_DATA UnitData
var SPELL_DATA SpellData
var SUMMONER_SPELL_DATA SpellData
var resourcepath string

func init() {
	resourcepath, _ = filepath.Abs("../hook-golang/src/resources")
}

func LoadUnitData() {
	file, err := ioutil.ReadFile(resourcepath + "/UnitData.json")

	if err != nil {
		fmt.Println("Error in ReadFile UNIT_DATA ", err)
	}
	data := UnitData{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println("Error in UnmarShal UNIT_DATA ", err)
	}
	UNIT_DATA = data
}

func SpelltData() {
	file, err := ioutil.ReadFile(resourcepath + "/SpellData.json")

	if err != nil {
		fmt.Println("Error in ReadFile SPELL_DATA ", err)
	}
	data := SpellData{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println("Error in UnmarShal SPELL_DATA ", err)
	}
	SPELL_DATA = data
}

func SummonerSpellData() {
	file, err := ioutil.ReadFile(resourcepath + "/SummonerSpellData.json")

	if err != nil {
		fmt.Println("Error in ReadFile SUMMONER_SPELL_DATA ", err)
	}
	data := SpellData{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println("Error in UnmarShal SUMMONER_SPELL_DATA ", err)
	}
	SUMMONER_SPELL_DATA = data
}
