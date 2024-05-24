package commands

import (
	"github.com/vonuki/textgame/internal/locations"
	"github.com/vonuki/textgame/internal/player"
)

type Command func(wr *locations.WorldHandler, pl player.Player, args ...string) string

type ComKeeper struct {
	commands map[string]Command
}

func NewCommandKeeper() *ComKeeper {
	cmds := make(map[string]Command)

	return &ComKeeper{
		commands: cmds,
	}
}

func (c *ComKeeper) GetCommandNames() string {
	if len(c.commands) == 0 {
		return "Нет доступных команд"
	}
	res := "Список доступных команд:"
	for name, _ := range c.commands {
		res = res + " " + name
	}
	return res
}

func (c *ComKeeper) AddCommand(name string, cmd Command) {
	c.commands[name] = cmd
}

func (c *ComKeeper) DoCommandByName(name string, wr *locations.WorldHandler, pl player.Player, args ...string) string {
	cmd, ok := c.commands[name]
	if !ok {
		return "неизвестная команда"
	}
	return cmd(wr, pl, args...)
}
