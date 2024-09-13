package item

import (
	condition "skit_ttrpy_backend/datatypes/conditions"
	"skit_ttrpy_backend/datatypes/gameplay/action/ability"
	itemcomponent "skit_ttrpy_backend/datatypes/item/item_component"
)

type EquipmentType int

const (
	Amulet = iota
	Armour
	Boots
	Cape
	Clothing
	Headware
	Ring
	Weapon
	Miscellaneous
)

type Equipment struct {
	ConditionComponent   itemcomponent.ItemConditionComponent
	DerivedStatComponent itemcomponent.DerivedStatComponent
	AbilityComponent     itemcomponent.ItemAbilityComponent
	EquipmentType        EquipmentType
	Name                 string
	Description          string
}

func GetEquipmentType(equipment *Equipment) EquipmentType {
	return equipment.EquipmentType
}

func GetAbilityComponent(equipment *Equipment) itemcomponent.ItemAbilityComponent {
	return equipment.AbilityComponent
}

func GetDerivedStatComponent(equipment *Equipment) itemcomponent.DerivedStatComponent {
	return equipment.DerivedStatComponent
}

func GetConditionComponent(equipment *Equipment) itemcomponent.ItemConditionComponent {
	return equipment.ConditionComponent
}

func GetNone(equipmentType EquipmentType) Equipment {
	return Equipment{
		ConditionComponent:   itemcomponent.NewItemConditionComponent(make([]condition.Condition, 10)),
		DerivedStatComponent: itemcomponent.NewDerivedStatComponent(nil, nil, nil, nil, nil),
		EquipmentType:        equipmentType,
		AbilityComponent:     itemcomponent.NewItemAbilityComponent(make([]ability.Ability, 10)),
		Name:                 "None",
		Description:          "None",
	}
}
