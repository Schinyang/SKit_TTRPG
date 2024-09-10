package character

import "fmt"

type Character struct {
	BaseStats    CharacterStats
	CurrentStats CharacterStats
	Level        int
	Conditions   []Condition
	Bulwark int
	Evasion int
	MainHandAttack int
	OffHandAttack int 
	RangedAttack int
	
}

func (character *Character) ApplyCondition(condition *Condition){
	condition.Apply(character)
	character.Conditions = append(character.Conditions, condition) 
}

func (character *Character) RemoveCondition(condition *Condition){
	condition.Remove(character)
	for i, eachCond in range(character.Conditions){
		if(condition.Name() == eachCond.Name()){
			character.Conditions = append(character.Conditions[:i], character.Conditions[:i+1]...)
			break
		}
	}

}

func (character *Character) GetCurentStats() CharacterStats {
	return CurrentStats
}

func (character *Character) GetBaseStats() CharacterStats {
	return BaseStats
}

func (character *Character) GainStats(stat Stat, increase int){
	BaseStats.AddStat(stat, increase)
}

func (character *Character) GainLevel() {
	Level += 1
}
