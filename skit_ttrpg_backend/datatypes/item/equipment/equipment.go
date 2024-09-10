package equipment

type Slot int
const {
	HeadSlot = iota
	CapeSlot
	ClothingSlot
	ArmourSlot
	AmuletSlot
	BootSlot
	RingSlot
} 
type equipment interface(){
	Equip(*Character)
	Unequip(*Character)
	Description() string
	getSlot() Slot
}