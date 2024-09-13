package stats

type Stat int

const (
	Strength = iota
	Dexterity
	Constitution
	Intelligence
	Mind
	Charisma
	Soul
)

var StatNames = map[Stat]string{
	Strength:     "Strength",
	Dexterity:    "Dexterity",
	Constitution: "Constitution",
	Intelligence: "Intelligence",
	Mind:         "Mind",
	Soul:         "Soul",
}

type CharacterStats struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Mind         int
	Charisma     int
	Soul         int
}

func NewCharacterStats(strength, dexterity, constitution, intelligence, charisma, mind, soul int) CharacterStats {
	return CharacterStats{
		Strength:     strength,
		Dexterity:    dexterity,
		Constitution: constitution,
		Intelligence: intelligence,
		Mind:         mind,
		Charisma:     charisma,
		Soul:         soul,
	}
}

func (cs CharacterStats) getStat(stat Stat) int {
	switch stat {
	case Strength:
		return cs.Strength
	case Dexterity:
		return cs.Dexterity
	case Constitution:
		return cs.Constitution
	case Intelligence:
		return cs.Intelligence
	case Mind:
		return cs.Mind
	case Charisma:
		return cs.Charisma
	case Soul:
		return cs.Soul
	default:
		return 0
	}
}

func (cs CharacterStats) AddStat(stat Stat, increase int) {
	switch stat {
	case Strength:
		cs.Strength += increase
	case Dexterity:
		cs.Dexterity += increase
	case Constitution:
		cs.Constitution += increase
	case Intelligence:
		cs.Intelligence += increase
	case Mind:
		cs.Mind += increase
	case Charisma:
		cs.Charisma += increase
	case Soul:
		cs.Soul += increase
	}
}
