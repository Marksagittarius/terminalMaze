package maze

import (
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/player"
)

type RoomMiddleware func(RoomPrototype, *player.Player, *logger.Logger) 

type DoorMiddleware func(DoorPrototype, *player.Player, *logger.Logger)

type MazeMiddleware func(*Maze, *player.Player, *logger.Logger)

func DefaultRoomMiddleware(RoomMiddleware, *player.Player, *logger.Logger) {

}

func DefaultDoorMiddleware(DoorPrototype, *player.Player, *logger.Logger) {

}

func DefaultMazeMiddleware(*Maze, *player.Player, *logger.Logger) {
	
}

func CombineRoomMiddlewares(middlewares... RoomMiddleware) RoomMiddleware {
	return func (room RoomPrototype, player *player.Player, logger *logger.Logger) {
		for _, handler := range middlewares {
			handler(room, player, logger)
		}
	}
}

func CombineDoorMiddlewares(middlewares... DoorMiddleware) DoorMiddleware {
	return func(door DoorPrototype, player *player.Player, logger *logger.Logger) {
		for _, handler := range middlewares {
			handler(door, player, logger)
		}
	}
}

func CombineMazeMiddlewares(middlewares... MazeMiddleware) MazeMiddleware {
	return func(maze *Maze, player *player.Player, logger *logger.Logger) {
		for _, handler := range middlewares {
			handler(maze, player, logger)
		}
	}
}