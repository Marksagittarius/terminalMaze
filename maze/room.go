package maze

import (
	"fmt"
	"reflect"
	"strconv"
	"github.com/Marksagittarius/terminalMaze/factory"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/monster"
	"github.com/Marksagittarius/terminalMaze/player"
)

type Room struct {
	roomId    string
	roomType  string
	neighbors [4]DoorPrototype
	preloads  []RoomMiddleware
	callbacks []RoomMiddleware
	monsters  []monster.MonsterPrototype
	visited   bool
}

func (room *Room) GetRoomId() string {
	return room.roomId
}

func (room *Room) GetRoomType() string {
	return room.roomType
}

func (room *Room) EnterDoor(direction Direction, player *player.Player, logger *logger.Logger) string {
	if logger.IsDead() || logger.IsOver() {
		return ""
	}

	currentDoor := room.neighbors[direction]

	if currentDoor == nil {
		logger.Println("[Door] There is no door in that direction.")
		return ""
	}

	currentDoor.ApplyPreloads(player, logger)
	currentDoor.Visit()
	nextRoom, ok := currentDoor.GetNextRoom(room.roomId)
	currentDoor.ApplyCallbacks(player, logger)

	if !ok {
		logger.Println("[Door] It must be illusion because there is actually no door.")
		return ""
	}

	if nextRoom == room.roomId {
		logger.Println("[Door] The door cannot be opened.")
		return ""
	}

	logger.Println("[Room] Moving from [" + room.roomId + "] to [" + nextRoom + "].")
	return nextRoom
}

func (room *Room) GetClass() reflect.Type {
	return reflect.TypeOf(room)
}

func (room *Room) ToString() string {
	ans := fmt.Sprintf("%v", room)
	return ans
}

func (room *Room) Copy() factory.AbstractProduct {
	copiedRoomID := room.roomId
	copiedRoomType := room.roomType

	copiedNeighbors := [4]DoorPrototype{}
	copy(copiedNeighbors[:], room.neighbors[:])
	copiedPreloads := make([]RoomMiddleware, len(room.preloads))
	copy(copiedPreloads, room.preloads)
	copiedCallbacks := make([]RoomMiddleware, len(room.preloads))
	copy(copiedCallbacks, room.callbacks)

	return &Room{
		roomId: copiedRoomID,
		roomType: copiedRoomType,
		neighbors: copiedNeighbors,
		preloads: copiedPreloads,
		callbacks: copiedCallbacks,
	}
} 

func (room *Room) AppendPreloads(middlewares... RoomMiddleware) RoomPrototype {
	room.preloads = append(room.preloads, middlewares...)
	return room
}

func (room *Room) AppendCallbacks(middlewares... RoomMiddleware) RoomPrototype {
	room.callbacks = append(room.callbacks, middlewares...)
	return room
}

func (room *Room) ClearPreloads() RoomPrototype {
	room.callbacks = make([]RoomMiddleware, 0)
	return room
}

func (room *Room) ClearCallbacks() RoomPrototype {
	room.preloads = make([]RoomMiddleware, 0) 
	return room
}

func (room *Room) SetNeighbor(direction Direction, door DoorPrototype) RoomPrototype {
	if direction >= 4 || direction < 0 {
		return room
	}

	room.neighbors[direction] = door
	return room
}

func (room *Room) ApplyPreloads(player *player.Player, logger *logger.Logger) {
	if logger.IsDead() || logger.IsOver() {
		return
	}

	for _, handler := range room.preloads {
		handler(room, player, logger)
	}
}

func (room *Room) ApplyCallbacks(player *player.Player, logger *logger.Logger) {
	if logger.IsDead() || logger.IsOver() {
		return
	}

	for _, handler := range room.callbacks {
		handler(room, player, logger)
	}
}

func (room *Room) HasVisited() bool {
	return room.visited
}

func (room *Room) Visit(player *player.Player, logger *logger.Logger, direction Direction) (RoomPrototype, Direction) {
	room.visited = true

	if room.GetMonstersNum() == 0 {
		return room, Null
	}

	for room.GetMonstersNum() > 0 {
		currentMonster := room.GetCurrentMonster()
		currentMonster.ApplyPreloads(player, logger)

		currentMonster.UpdateHeathPoint(-1 * player.GetAttackPoint())
		player.UpdateHeathPoint(-1 * currentMonster.GetAttackPoint())

		logger.Println("[Battle] <" + player.GetName() + "> gains " + strconv.Itoa(currentMonster.GetAttackPoint()) + " damages.")
		logger.Println("[Battle] <" + currentMonster.GetMonsterName() + "> gains " + strconv.Itoa(player.GetAttackPoint()) + " damages.")
		logger.Println("[Battle] <" + player.GetName() + "> has " + strconv.Itoa(player.GetHealthPoint()) + " HP left.")

		if player.GetHealthPoint() <= 0 || player.GetRationalPoint() <= 0 {
			logger.CondemnToDeath()
			return room, Null
		}

		if currentMonster.GetHealthPoint() == 0 {
			player.UpdateRationalPoint(1)
			currentMonster.ApplyCallbacks(player, logger)
			logger.Println("[Battle] <" + currentMonster.GetMonsterName() + "> disappears.")
			room.RemoveMonster()
			continue
		} else {
			currentMonster.ApplyCallbacks(player, logger)
			return room, GetReverseDirection(direction)
		}
	}

	return room, Null
}

func (room *Room) AppendMonsters(monsters... monster.MonsterPrototype) RoomPrototype {
	room.monsters = append(room.monsters, monsters...)
	return room
}

func (room *Room) GetCurrentMonster() monster.MonsterPrototype {
	if len(room.monsters) == 0 {
		return nil
	}

	return room.monsters[0]
}

func (room *Room) RemoveMonster() RoomPrototype {
	room.monsters = room.monsters[1:]
	return room
}

func (room *Room) ClearMonsters() RoomPrototype {
	room.monsters = make([]monster.MonsterPrototype, 0)
	return room
}

func (room *Room) GetMonstersNum() int {
	return len(room.monsters)
}












