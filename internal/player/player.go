package player

import (
	"github.com/vonuki/textgame/internal/items"
)

type Player interface {
	GetPlayerName() string
	GetCurrentLocation() string
	SetCurrentLocation(name string)
	items.ItemStorage
}

// DefaultPlayer - имплементирует интерфейсы Игрока и Хранилища
type DefaultPlayer struct {
	name            string
	currentLocation string
	items.DefaultStorage
}

func NewPlayer(name string, storageName string, locationName string) *DefaultPlayer {
	it := make([]items.Item, 0)

	return &DefaultPlayer{
		name:            name,
		currentLocation: locationName,
		DefaultStorage: items.DefaultStorage{
			StorageName: storageName,
			Items:       it,
		},
	}
}

func (p *DefaultPlayer) GetPlayerName() string {
	return p.name
}
func (p *DefaultPlayer) GetCurrentLocation() string {
	return p.currentLocation
}
func (p *DefaultPlayer) SetCurrentLocation(name string) {
	p.currentLocation = name
}
