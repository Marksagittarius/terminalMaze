package maze

import (
	"fmt"
	"reflect"
	"github.com/Marksagittarius/terminalMaze/item"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/monster"
	"github.com/Marksagittarius/terminalMaze/player"
	"github.com/Marksagittarius/terminalMaze/terminal"
)

type Maze struct {
	roomsList     map[string]RoomPrototype
	doorsList     map[string]DoorPrototype
	logger        *logger.Logger
	player        *player.Player
	controller	  terminal.TerminalController
	startRoomId   string
	currentRoomId string
	preloads      []MazeMiddleware
	callbacks     []MazeMiddleware
	itemsList     map[string]item.ItemPrototype
	monstersList  map[string]monster.MonsterPrototype
}

func NewMaze(startRoomId string) *Maze {
	newLogger := logger.NewLogger()
	return &Maze{
		roomsList: make(map[string]RoomPrototype),
		doorsList: make(map[string]DoorPrototype),
		player: nil,
		logger: newLogger,
		controller: terminal.NewDefaultTerminalController(newLogger),
		startRoomId: startRoomId,
		currentRoomId: startRoomId,
		preloads: make([]MazeMiddleware, 0),
		callbacks: make([]MazeMiddleware, 0),
		itemsList: make(map[string]item.ItemPrototype),
		monstersList: make(map[string]monster.MonsterPrototype),
	}
}

func (maze *Maze) GetClass() reflect.Type {
	return reflect.TypeOf(maze)
}

func (maze *Maze) ToString() string {
	ans := fmt.Sprintf("%v", maze)
	return ans
}

func (maze *Maze) Copy() *Maze {
	return nil
}

func (maze *Maze) AppendPreloads(middlewares ...MazeMiddleware) *Maze {
	maze.preloads = append(maze.preloads, middlewares...)
	return maze
}

func (maze *Maze) AppendCallbacks(middlewares ...MazeMiddleware) *Maze {
	maze.callbacks = append(maze.callbacks, middlewares...)
	return maze
}

func (maze *Maze) ClearPreloads() *Maze {
	maze.preloads = make([]MazeMiddleware, 0)
	return maze
}

func (maze *Maze) ClearCallbacks() *Maze {
	maze.callbacks = make([]MazeMiddleware, 0)
	return maze
}

func (maze *Maze) AppendRooms(rooms ...RoomPrototype) *Maze {
	for _, room := range rooms {
		maze.roomsList[room.GetRoomId()] = room
	}

	return maze
}

func (maze *Maze) AppendDoors(doors ...DoorPrototype) *Maze {
	for _, door := range doors {
		maze.doorsList[door.GetDoorId()] = door
	}

	return maze
}

func (maze *Maze) SetPlayer(player *player.Player) *Maze {
	maze.player = player
	return maze
}

func (maze *Maze) SetStartRoom(roomId string) *Maze {
	if _, ok := maze.roomsList[roomId]; !ok {
		return maze
	}

	maze.startRoomId = roomId
	return maze
}

func (maze *Maze) AppendItems(items ...item.ItemPrototype) *Maze {
	for _, item := range items {
		maze.itemsList[item.GetItemName()] = item
	}

	return maze
}

func (maze *Maze) GetCurrentRoomId() string {
	return maze.currentRoomId
}

func (maze *Maze) ReplaceLogger(logger *logger.Logger) *Maze {
	maze.logger = logger
	return maze
}

func (maze *Maze) AppendMonsters(monsters... monster.MonsterPrototype) *Maze {
	for _, monster := range monsters {
		maze.monstersList[monster.GetMonsterName()] = monster
	}

	return maze
}

func (maze *Maze) ApplyPreloads(player *player.Player, logger *logger.Logger) {
	if maze.logger.IsDead() || maze.logger.IsOver() {
		return
	}

	for _, handler := range maze.preloads {
		handler(maze, player, logger)
	}
}

func (maze *Maze) ApplyCallbacks(player *player.Player, logger *logger.Logger) {
	for _, handler := range maze.callbacks {
		handler(maze, player, logger)
	}
}

func (maze *Maze) SetMonsters(monsters []string, room string) *Maze {
	if _, hasRoom := maze.roomsList[room]; !hasRoom {
		return maze
	}

	currentRoom := maze.roomsList[room]
	for _, monsterName := range monsters {
		if monsterInstance, ok := maze.monstersList[monsterName]; ok {
			currentRoom.AppendMonsters(monsterInstance.Copy())
		}
	}

	return maze
}

func (maze *Maze) Start() {
	if maze.player == nil {
		maze.logger.Println("[Error] No player loaded.")
		return 
	}

	if len(maze.roomsList) == 0 {
		maze.logger.Println("[Error] No Room loaded.")
		return
	}

	if maze.logger.IsDead() {
		maze.logger.Println("[Error] No player alive.")
		return
	}

	maze.ApplyPreloads(maze.player, maze.logger)
	for !maze.logger.IsOver() {

		if maze.logger.IsDead() {
			maze.logger.GameOver()
			break
		}
		
		currentRoom := maze.roomsList[maze.currentRoomId]

		maze.logger.Println("[Room] Room: " + currentRoom.GetRoomId())
		maze.player.ApplyPreloads(maze.logger)
		currentRoom.ApplyPreloads(maze.player, maze.logger)
		_, newDirection := currentRoom.Visit(maze.player, maze.logger, Direction(maze.logger.GetLastOperation()))

		if maze.logger.IsDead() {
			maze.logger.GameOver()
			break
		}

		var direction Direction
		if newDirection != Null {
			direction = newDirection
		} else {
			commandLeft, commandRight := "", ""
			for {
				maze.logger.Println("[Menu] What will <" + maze.player.GetName() + "> do ?")
				maze.logger.Println("[Command] move to {direction} : Move to another room.")
				maze.logger.Println("[Command] display items : Display all the items.")
				maze.logger.Println("[Command] use {itemName} : Use the item.")
				commandLeft, commandRight = maze.controller.ReadMenuCommand()
				if commandLeft == "display" && commandRight == "items" {
					maze.player.DisplayItems(maze.logger)
					continue
				} else if commandLeft == "use" {
					if _, ok := maze.itemsList[commandRight]; ok {
						itemName := maze.player.UseItem(commandRight)
						if itemName != "" {
							if maze.player.GetRationalPoint() >= maze.itemsList[commandRight].GetRationalPoint() {
								maze.itemsList[commandRight].UseEffect(maze.player, maze.logger)
							}
						}
						continue
					} else {
						maze.logger.Println("[Warning] No such item.")
					}
				} else if commandLeft == "move" {
					direction = Direction(maze.controller.ReadDirection())
					break
				}
			}
			
		}

		nextRoom := currentRoom.EnterDoor(Direction(direction), maze.player, maze.logger)

		if nextRoom == "" {
			maze.logger.Println("[Room] What a waste of time.")
			continue
		}

		maze.logger.AppendOperation(int(direction))
		currentRoom.ApplyCallbacks(maze.player, maze.logger)
		maze.player.ApplyCallbacks(maze.logger)
		maze.currentRoomId = nextRoom
	}
	maze.ApplyCallbacks(maze.player, maze.logger)
}
