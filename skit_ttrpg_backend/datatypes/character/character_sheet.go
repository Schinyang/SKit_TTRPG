package character

type Character struct {
	BaseStats    CharacterStats
	CurrentStats CharacterStats
	Level        int
	Conditions   []Condition
}
