package maze

import (
	"github.com/Marksagittarius/terminalMaze/factory"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/monster"
	"github.com/Marksagittarius/terminalMaze/player"
)

type RoomPrototype interface {
	factory.AbstractProduct
	GetRoomId() string
	GetRoomType() string
	EnterDoor(Direction, *player.Player, *logger.Logger) string
	AppendPreloads(...RoomMiddleware) RoomPrototype
	AppendCallbacks(...RoomMiddleware) RoomPrototype
	ClearPreloads() RoomPrototype
	ClearCallbacks() RoomPrototype
	SetNeighbor(Direction, DoorPrototype) RoomPrototype
	ApplyPreloads(*player.Player, *logger.Logger)
	ApplyCallbacks(*player.Player, *logger.Logger)
	HasVisited() bool
	Visit(*player.Player, *logger.Logger, Direction) (RoomPrototype, Direction)
	AppendMonsters(...monster.MonsterPrototype) RoomPrototype
	GetCurrentMonster() monster.MonsterPrototype
	RemoveMonster() RoomPrototype
	ClearMonsters() RoomPrototype
	GetMonstersNum() int
}