package maze

func ConnectTwoRooms(room1 RoomPrototype, room2 RoomPrototype, door DoorPrototype, direction Direction) {
	door.Connect(room1.GetRoomId(), room2.GetRoomId())
	room1.SetNeighbor(direction, door)
	room2.SetNeighbor(GetReverseDirection(direction), door)
}
