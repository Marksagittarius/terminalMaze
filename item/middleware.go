package item

import (
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/player"
)

type ItemMiddleware func(*player.Player, *logger.Logger)

func DefaultItemMiddleware(*player.Player, *logger.Logger) {
	
}

func CombineItemMiddlewares(middlewares... ItemMiddleware) ItemMiddleware {
	return func(player *player.Player, logger *logger.Logger) {
		for _, handler := range middlewares {
			handler(player, logger)
		}
	}
}