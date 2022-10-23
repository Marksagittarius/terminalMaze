package main

import (
	"fmt"
	"github.com/Marksagittarius/terminalMaze/item"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/maze"
	"github.com/Marksagittarius/terminalMaze/monster"
	"github.com/Marksagittarius/terminalMaze/player"
	"github.com/Marksagittarius/terminalMaze/util"
)

func main() {
	roomFactory := maze.NewRoomFactory()
	doorFactory := maze.NewDoorFactory()
	monsterFactory := monster.NewMonsterFactory()
	itemFactory := item.NewItemFactory()

	startRoom := roomFactory.Produce(&maze.RoomProductConfig{
		RoomId: "Starter Room",
		RoomType: "Normal",
	})

	northRoom := roomFactory.Produce(&maze.RoomProductConfig{
		RoomId: "North Room",
		RoomType: "Normal",
	})

	northDoor := doorFactory.Produce(&maze.DoorProductConfig{
		DoorId: "North Door",
		DoorType: "Normal",
	})

	startRoom.AppendPreloads(func(room maze.RoomPrototype, player *player.Player, logger *logger.Logger) {
		logger.Println("<Room Preload>")
		if !room.HasVisited() {
			logger.Println("[Item]<" + player.GetName() + "> got 1 [Gehao-Medicine]." )
			player.AddItems("Gehao-Medicine", 1)
		}
	})

	startRoom.AppendCallbacks(func(_ maze.RoomPrototype, _ *player.Player, logger *logger.Logger) {
		logger.Println("<Room Callback>")
	})

	northDoor.AppendPreloads(func(_ maze.DoorPrototype, _ *player.Player, logger *logger.Logger) {
		logger.Println("<Door Preload>")
	})

	northDoor.AppendCallbacks(func(_ maze.DoorPrototype, _ *player.Player, logger *logger.Logger) {
		logger.Println("<Door Callback>")
	})

	northRoom.AppendCallbacks(func(_ maze.RoomPrototype, _ *player.Player, logger *logger.Logger) {
		logger.Println("<NorthRoom Callback>")
	})

	maze.ConnectTwoRooms(startRoom, northRoom, northDoor, maze.North)

	northMonster := monsterFactory.Produce(&monster.MonsterProductConfig{
		MonsterName: "Gehao Monster",
		InitHealthPoint: 2,
		InitAttackPoint: 1,
	})

	northMonster.AppendPreloads(func(monster monster.MonsterPrototype, _ *player.Player, logger *logger.Logger) {
		if monster.GetHealthPoint() == monster.GetInitHealthPoint() {
			logger.Println("[Monster]" + " <" + monster.GetMonsterName() + "> appears.")
			fmt.Println("<" + monster.GetMonsterName() + ">")
			fmt.Println(monster.ToString())
		}
	})

	northMonster.AppendCallbacks(func(monster monster.MonsterPrototype, _ *player.Player, logger *logger.Logger) {
		if monster.GetHealthPoint() == 0 {
			logger.Println("[Monster]" + " <" + monster.GetMonsterName() + "> turns into dust.")
			return
		}
	})

	defaultPlayer := player.NewPlayer(&player.PlayerConfig{
		Name: "Gehao",
		InitHealthPoint: 100,
		InitRationalPoint: 10,
		InitAttackPoint: 1,
	})

	defaultPlayer.AppendPreloads(func(_ *player.Player, logger *logger.Logger) {
		util.DividerPrinter(100)
		logger.Println("<Player Preload>")
	})

	defaultPlayer.AppendCallbacks(func(_ *player.Player, logger *logger.Logger) {
		logger.Println("<Player Callback>")
		util.DividerPrinter(100)
	})

	mazeDemo := maze.NewMaze("Starter Room")
	mazeDemo.SetPlayer(defaultPlayer)
	mazeDemo.AppendDoors(northDoor)
	mazeDemo.AppendRooms(startRoom, northRoom)

	mazeDemo.AppendPreloads(func(_ *maze.Maze, _ *player.Player, logger *logger.Logger) {
		logger.Println("<Maze Preload>")
	})

	mazeDemo.AppendCallbacks(func(_ *maze.Maze, _ *player.Player, logger *logger.Logger) {
		if logger.IsDead() {
			logger.Println("[The End] Rest In Peace.")
		}
		logger.Println("<Maze Callback>")
	})

	mazeDemo.AppendMonsters(northMonster)

	gehaoItem := itemFactory.Produce(&item.ItemProductConfig{
		ItemName: "Gehao-Medicine",
		RationalPoint: 1,
	})

	gehaoItem.SetEffect(func(player *player.Player, logger *logger.Logger) {
		logger.Println("[Item] Recover the player by adding 10 HP.")
		player.UpdateHeathPoint(10)
	})

	mazeDemo.SetMonsters([]string{"Gehao Monster", "Gehao Monster"}, "North Room")
	mazeDemo.AppendItems(gehaoItem)

	mazeDemo.Start()
}