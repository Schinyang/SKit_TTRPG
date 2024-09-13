package character

import "skit_ttrpy_backend/component"

type Character struct {
	BaseStats          component.StatComponent
	CurrentStats       component.StatComponent
	Level              int
	ConditionComponent []component.ConditionComponent
}
