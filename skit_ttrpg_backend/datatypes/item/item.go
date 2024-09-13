package item

type Item struct {
	Name        string
	Description string
}

func GetName(item *Item) string {
	return (item.Name)
}

func GetDescription(item Item) string {
	return (item.Description)
}
