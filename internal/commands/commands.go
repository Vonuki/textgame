package commands

import (
	"github.com/vonuki/textgame/internal/locations"
	"github.com/vonuki/textgame/internal/player"
	"strings"
)

// Пакет с реализацие различных команд

func LookAround(wr *locations.WorldHandler, pl player.Player, args ...string) string {
	loc := wr.GetLocationByName(pl.GetCurrentLocation())
	if loc == nil {
		return "нет такой локации"
	}

	resp := ""
	if loc.GetStorageName() != "" {
		if len(loc.GetItemNames()) > 0 {
			resp = loc.GetStorageName() + ": " + strings.Join(loc.GetItemNames(), ", ")
		} else {
			resp = "пустая комната"
		}
	} else {
		resp = loc.LookAroundAnswer()
	}
	resp = resp + "."

	locNames := loc.GetLinkedLocationsNames()
	if len(locNames) > 0 {
		resp = resp + " можно пройти - " + strings.Join(locNames, ", ")
	}

	return resp
}

func ComeIn(wr *locations.WorldHandler, pl player.Player, args ...string) string {
	loc := wr.GetLocationByName(pl.GetCurrentLocation())
	if loc == nil {
		return "нет такой локации"
	}

	resp := ""
	answ := loc.ComeInAnswer()
	if len(answ) > 0 {
		resp = resp + answ
	}
	resp = resp + "."

	locNames := loc.GetLinkedLocationsNames()
	if len(locNames) > 0 {
		resp = resp + " можно пройти - " + strings.Join(locNames, ", ")
	}

	return resp
}

func Move(wr *locations.WorldHandler, pl player.Player, args ...string) string {

	if len(args) != 1 {
		return "неизвестная команда"
	}
	dest := args[0]
	curLoc := wr.GetLocationByName(pl.GetCurrentLocation())
	destLoc := wr.GetLocationByName(dest)
	if curLoc == nil {
		return "нет такой локации " + pl.GetCurrentLocation()
	}
	if destLoc == nil {
		return "нет такой локации " + dest
	}

	if !curLoc.HasPathTo(dest) {
		return "нет пути в " + dest
	}

	pl.SetCurrentLocation(dest)

	return ComeIn(wr, pl, args...)
}

func Take(wr *locations.WorldHandler, pl player.Player, args ...string) string {

	if len(args) != 1 {
		return "неизвестная команда"
	}
	itemName := args[0]
	curLoc := wr.GetLocationByName(pl.GetCurrentLocation())

	item := curLoc.TakeItem(itemName)
	if item == nil {
		return "нет такого"
	}
	pl.PutItem(item)

	return "предмет добавлен в " + pl.GetStorageName() + ": " + item.GetItemName()
}
