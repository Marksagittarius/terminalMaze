package admin

import (
	"sync"
	"github.com/Marksagittarius/terminalMaze/maze"
)

type MazeController interface {
	AddMaze(*maze.Maze) MazeController
	Start()
}

type mazeController struct {
	mazes []*maze.Maze
}

func (mc *mazeController) AddMaze(maze *maze.Maze) MazeController {
	mc.mazes = append(mc.mazes, maze)
	return mc
}

func (mc *mazeController) Start() {
	for _, maze := range mc.mazes {
		maze.Start()
	}
}

var (
	instance *mazeController
	once sync.Once
)

func GetMazeControllerInstance() MazeController {
	once.Do(func() {
		instance = &mazeController{
			mazes: make([]*maze.Maze, 0),
		}
	})

	return instance
}