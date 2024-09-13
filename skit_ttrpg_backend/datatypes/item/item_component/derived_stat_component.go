package itemcomponent

type DerivedStatComponent struct {
	MainHandAttackValue int
	OffHandAttackValue  int
	CriticalHitBoundary int
	Evasion             int
	Bulwark             int
}

func NewDerivedStatComponent(mhav, ofav, chb, eva, bul *int) DerivedStatComponent {
	return DerivedStatComponent{
		MainHandAttackValue: derefOrDefault(mhav, 0),
		OffHandAttackValue:  derefOrDefault(ofav, 0),
		CriticalHitBoundary: derefOrDefault(chb, 0),
		Evasion:             derefOrDefault(eva, 0),
		Bulwark:             derefOrDefault(bul, 0),
	}

}

func derefOrDefault(val *int, defaultVal int) int {
	if val != nil {
		return *val
	}
	return defaultVal
}
