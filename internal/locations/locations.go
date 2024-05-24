package locations

import "github.com/vonuki/textgame/internal/items"

type Location interface {
	GetLocationName() string
	GetLinkedLocationsNames() []string
	AddLinkedLocation(name string)
	LookAroundAnswer() string
	ComeInAnswer() string
	HasPathTo(dest string) bool
	items.ItemStorage
}

// DefaultLocation - имплементирует интерфейсы Локации и Хранилища
type DefaultLocation struct {
	name  string
	links []string
	items.DefaultStorage
	lookAroundAnswer string
	comeInAnswer     string
}

func NewLocation(name, comeInAnswer, storageName, lookAroundAnswer string) *DefaultLocation {
	it := make([]items.Item, 0)
	lk := make([]string, 0)

	return &DefaultLocation{
		name:         name,
		comeInAnswer: comeInAnswer,
		links:        lk,
		DefaultStorage: items.DefaultStorage{
			StorageName: storageName,
			Items:       it,
		},
		lookAroundAnswer: lookAroundAnswer,
	}
}

func (l *DefaultLocation) GetLocationName() string {
	return l.name
}

func (l *DefaultLocation) GetLinkedLocationsNames() []string {
	return l.links
}

func (l *DefaultLocation) AddLinkedLocation(name string) {
	l.links = append(l.links, name)
}

func (l *DefaultLocation) LookAroundAnswer() string {
	return l.lookAroundAnswer
}

func (l *DefaultLocation) ComeInAnswer() string {
	return l.comeInAnswer
}

func (l *DefaultLocation) HasPathTo(dest string) bool {
	for i := 0; i < len(l.links); i++ {
		if dest == l.links[i] {
			return true
		}
	}
	return false
}
