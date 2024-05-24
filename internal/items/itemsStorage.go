package items

type ItemStorage interface {
	GetStorageName() string
	GetItems() []Item
	GetItemNames() []string
	TakeItem(name string) Item
	PutItem(Item)
}

type DefaultStorage struct {
	StorageName string
	Items       []Item
}

func (s *DefaultStorage) GetStorageName() string {
	return s.StorageName
}

func (s *DefaultStorage) GetItems() []Item {
	return s.Items
}

func (s *DefaultStorage) GetItemNames() []string {
	names := make([]string, 0)
	for _, item := range s.Items {
		names = append(names, item.GetItemName())
	}
	return names
}

func (s *DefaultStorage) TakeItem(name string) Item {
	for i, item := range s.Items {
		if item.GetItemName() == name {
			res := s.Items[i]

			copy(s.Items[i:], s.Items[i+1:])
			s.Items[len(s.Items)-1] = nil
			s.Items = s.Items[:len(s.Items)-1]
			return res
		}
	}
	return nil
}

func (s *DefaultStorage) PutItem(item Item) {
	s.Items = append(s.Items, item)
}
