package character

type Condition interface {
	Name() string
	Apply(*Character)
	Remove(*Character)
	Description() string
}