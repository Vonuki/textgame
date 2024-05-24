package items

type Item interface {
	GetItemName() string
}

type DefaultItem struct {
	name string
}

func NewItem(name string) *DefaultItem {
	return &DefaultItem{name: name}
}

func (i *DefaultItem) GetItemName() string {
	return i.name
}
