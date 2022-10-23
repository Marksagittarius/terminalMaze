package player

import "github.com/Marksagittarius/terminalMaze/logger"

type PlayerMiddleware func(*Player, *logger.Logger)

func DefaultPlayerMiddleware(_ *Player, _ *logger.Logger) {

}

func CombinePlayerMiddlewares(middlewares... PlayerMiddleware) PlayerMiddleware {
	return func(player *Player, logger *logger.Logger) {
		for _, handler := range middlewares {
			handler(player, logger)
		}
	}
}