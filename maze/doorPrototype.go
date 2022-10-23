package maze

import (
	"github.com/Marksagittarius/terminalMaze/factory"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/player"
)

type DoorPrototype interface {
	factory.AbstractProduct
	GetDoorId() string
	GetDoorType() string
	AppendPreloads(...DoorMiddleware) DoorPrototype
	AppendCallbacks(...DoorMiddleware) DoorPrototype
	ClearPreloads() DoorPrototype
	ClearCallbacks() DoorPrototype
	GetNextRoom(string) (string, bool)
	Connect(string, string) DoorPrototype
	ApplyPreloads(*player.Player, *logger.Logger)
	ApplyCallbacks(*player.Player, *logger.Logger)
	HasVisited() bool
	Visit() DoorPrototype
}