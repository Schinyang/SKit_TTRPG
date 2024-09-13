package component

import "skit_ttrpy_backend/datatypes/character/stats"

type StatComponent struct {
	baseStats    stats.CharacterStats
	currentStats stats.CharacterStats
}
