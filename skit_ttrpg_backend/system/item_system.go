package system

import (
	"skit_ttrpy_backend/component"
	"skit_ttrpy_backend/datatypes/item/equipment"
)

type ItemSystem struct{}

func (itemSystem *ItemSystem) EquipEquipment(derivedStatsComponent component.DerivedStatsComponent, item equipment.Equipment)
