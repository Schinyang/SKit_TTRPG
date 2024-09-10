package character

type SkillProficiency int

const{
	Insight = iota
	Society
	Nature
	Animal Handling
	Persuasion
	Deception
	Performance
	Intimidation
	Arcane
	MonsterAnatomy
	Philosophy
	Athletics
	Acrobatics
	MonsterAnatomy
	Alchol Resistance
	Technology
}

type CombatProficiency int

const {
	LightArmour
	MediumArmour
	HeavyArmour
	MartialWeapons
	SimpleWeapons
	Will
	Reflex
	Fortitude
}

type ProficiencyType int

const {
	Untrained = iota
	Trained
	Expert
	Master
	Legendary
}