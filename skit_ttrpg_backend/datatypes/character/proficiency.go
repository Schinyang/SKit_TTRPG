package character

type SkillProficiency int

const (
	Insight = iota
	Society
	Nature
	AnimalHandling
	Persuasion
	Deception
	Performance
	Intimidation
	Arcane
	MonsterAnatomy
	Philosophy
	Athletics
	Acrobatics
	AlcholResistance
	Technology
)

type CombatProficiency int

const (
	LightArmour = iota
	MediumArmour
	HeavyArmour
	MartialWeapons
	SimpleWeapons
	Will
	Reflex
	Fortitude
)

type ProficiencyType int

const (
	Untrained = iota
	Trained
	Expert
	Master
	Legendary
)
