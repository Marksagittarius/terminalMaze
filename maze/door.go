package maze

import (
	"fmt"
	"reflect"
	"github.com/Marksagittarius/terminalMaze/factory"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/player"
)

type Door struct {
	connection map[string]string
	doorType   string
	doorId     string
	preloads   []DoorMiddleware
	callbacks  []DoorMiddleware
	visited    bool
}

func (door *Door) GetDoorId() string {
	return door.doorId
}

func (door *Door) GetDoorType() string {
	return door.doorType
}

func (door *Door) GetClass() reflect.Type {
	return reflect.TypeOf(door)
}

func (door *Door) ToString() string {
	ans := fmt.Sprintf("%v", door)
	return ans
}

func (door *Door) AppendPreloads(middlewares... DoorMiddleware) DoorPrototype {
	door.preloads = append(door.preloads, middlewares...)
	return door
}

func (door *Door) AppendCallbacks(middlewares... DoorMiddleware) DoorPrototype {
	door.callbacks = append(door.callbacks, middlewares...)
	return door
}

func (door *Door) ClearPreloads() DoorPrototype {
	door.preloads = make([]DoorMiddleware, 0)
	return door
}

func (door *Door) ClearCallbacks() DoorPrototype {
	door.callbacks = make([]DoorMiddleware, 0)
	return door
}

func (door *Door) Copy() factory.AbstractProduct {
	copiedConnection := make(map[string]string)
	for key, value := range door.connection {
		copiedConnection[key] = value
	}

	copiedDoorId := door.doorId
	copiedDoorType := door.doorType
	copiedPreloads := make([]DoorMiddleware, len(door.preloads))
	copy(copiedPreloads, door.preloads)
	copiedCallbacks := make([]DoorMiddleware, len(door.callbacks))
	copy(copiedCallbacks, door.callbacks)

	return &Door{
		doorId: copiedDoorId,
		doorType: copiedDoorType,
		connection: copiedConnection,
		preloads: copiedPreloads,
		callbacks: copiedCallbacks,
	}
}

func (door *Door) GetNextRoom(roomId string) (string, bool) {
	next, ok := door.connection[roomId]
	return next, ok
} 

func (door *Door) Connect(left, right string) DoorPrototype {
	door.connection = make(map[string]string)
	door.connection[left] = right
	door.connection[right] = left
	return door
}

func (door *Door) ApplyPreloads(player *player.Player, logger *logger.Logger) {
	if logger.IsDead() || logger.IsOver() {
		return
	}

	for _, handler := range door.preloads {
		handler(door, player, logger)
	}
}

func (door *Door) ApplyCallbacks(player *player.Player, logger *logger.Logger) {
	if logger.IsDead() || logger.IsOver() {
		return
	}
	
	for _, handler := range door.callbacks {
		handler(door, player, logger)
	}
}

func (door *Door) HasVisited() bool {
	return door.visited
}

func (door *Door) Visit() DoorPrototype {
	door.visited = true
	return door
}


