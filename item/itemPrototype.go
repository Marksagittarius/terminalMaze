package item

import (
	"github.com/Marksagittarius/terminalMaze/factory"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/player"
)

type ItemPrototype interface {
	factory.AbstractProduct
	GetItemName() string
	GetRationalPoint() int
	SetEffect(ItemMiddleware) ItemPrototype
	UseEffect(*player.Player, *logger.Logger)
}