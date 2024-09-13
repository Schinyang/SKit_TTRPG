package itemcomponent

import "skit_ttrpy_backend/datatypes/gameplay/action/ability"

type ItemAbilityComponent struct {
	Ability []ability.Ability
}

func NewItemAbilityComponent(abilities []ability.Ability) ItemAbilityComponent {
	return (ItemAbilityComponent{Ability: abilities})
}
