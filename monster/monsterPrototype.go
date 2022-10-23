package monster

import (
	"github.com/Marksagittarius/terminalMaze/factory"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/player"
)

type MonsterPrototype interface {
	factory.AbstractProduct
	Copy() MonsterPrototype
	GetMonsterName() string
	GetHealthPoint() int
	GetAttackPoint() int
	GetInitHealthPoint() int
	GetInitAttackPoint() int
	AppendPreloads(...MonsterMiddleware) MonsterPrototype
	AppendCallbacks(...MonsterMiddleware) MonsterPrototype
	ClearPreloads() MonsterPrototype
	ClearCallbacks() MonsterPrototype
	UpdateHeathPoint(int) (MonsterPrototype, bool)
	UpdateAttackPoint(int) MonsterPrototype
	ApplyPreloads(*player.Player, *logger.Logger)
	ApplyCallbacks(*player.Player, *logger.Logger)
}

