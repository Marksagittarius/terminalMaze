package maze

import "github.com/Marksagittarius/terminalMaze/monster"

type RoomFactory struct {
}

type RoomProductConfig struct {
	RoomId   string
	RoomType string
}

func (rf *RoomFactory) Produce(config *RoomProductConfig) RoomPrototype {
	return &Room{
		roomId:    config.RoomId,
		roomType:  config.RoomType,
		neighbors: [4]DoorPrototype{},
		preloads:  make([]RoomMiddleware, 0),
		callbacks: make([]RoomMiddleware, 0),
		monsters:  make([]monster.MonsterPrototype, 0),
		visited:   false,
	}
}

func NewRoomFactory() *RoomFactory {
	return &RoomFactory{}
}
