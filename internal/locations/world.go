package locations

type WorldHandler struct {
	Locations map[string]Location
}

func NewWorldHandler() *WorldHandler {
	lk := make(map[string]Location)

	return &WorldHandler{
		Locations: lk,
	}
}

func (l *WorldHandler) AddLocation(location Location) {
	l.Locations[location.GetLocationName()] = location
}

func (l *WorldHandler) GetLocationByName(name string) Location {
	loc, ok := l.Locations[name]
	if !ok {
		return nil
	}
	return loc
}
