package component

import condition "skit_ttrpy_backend/datatypes/conditions"

type ConditionComponent struct {
	conditions []condition.Condition
}

func NewConditionComponent(initalConditions []condition.Condition) *ConditionComponent {
	return &ConditionComponent{conditions: initalConditions}
}
