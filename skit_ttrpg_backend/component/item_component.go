package component

import (
	"skit_ttrpy_backend/datatypes/item"
)

type ItemComponent struct {
	Items        []item.Item
	EquipmentMap map[item.EquipmentType]interface{}
}

func NewItemComponent() ItemComponent {
	rings := make([]item.Equipment, 2)
	rings[0] = item.GetNone(item.Ring)
	rings[1] = item.GetNone(item.Ring)
	equipmentMap := map[item.EquipmentType]interface{}{
		item.Amulet:   item.GetNone(item.Amulet),
		item.Armour:   item.GetNone(item.Armour),
		item.Boots:    item.GetNone(item.Boots),
		item.Cape:     item.GetNone(item.Boots),
		item.Clothing: item.GetNone(item.Clothing),
		item.Headware: item.GetNone(item.Headware),
		item.Ring:     rings,
	}
	return ItemComponent{
		Items:        make([]item.Item, 0, 10),
		EquipmentMap: equipmentMap,
	}
}
