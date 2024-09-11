package equipment

import "skit_ttrpy_backend/datatypes/character"

type Slot int

const (
	HeadSlot = iota
	CapeSlot
	ClothingSlot
	ArmourSlot
	AmuletSlot
	BootSlot
	RingSlot
)

type equipment interface {
	Equip(character *character.Character)
	Unequip(*character.Character)
	Description() string
	getSlot() Slot
}
