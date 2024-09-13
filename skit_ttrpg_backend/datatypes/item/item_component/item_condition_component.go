package itemcomponent

import condition "skit_ttrpy_backend/datatypes/conditions"

type ItemConditionComponent struct {
	AppliedConditions []condition.Condition
}

func NewItemConditionComponent(conditions []condition.Condition) ItemConditionComponent {
	return ItemConditionComponent{
		AppliedConditions: conditions,
	}
}
