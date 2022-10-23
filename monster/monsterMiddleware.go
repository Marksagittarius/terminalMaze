package monster

import (
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/player"
)

type MonsterMiddleware func(MonsterPrototype, *player.Player, *logger.Logger)

func DefaultMonsterMiddleware(MonsterPrototype, *player.Player, *logger.Logger) {

}

func CombineMonsterMiddlewares(middlewares... MonsterMiddleware) MonsterMiddleware {
	return func(monster MonsterPrototype, player *player.Player, logger *logger.Logger) {
		for _, handler := range middlewares {
			handler(monster, player, logger)
		}
	}
}