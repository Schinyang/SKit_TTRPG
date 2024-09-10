package character

type Condition interface {
	apply(c *Character)
}
