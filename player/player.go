package player

import (
	"fmt"
	"reflect"
	"strconv"
	"github.com/Marksagittarius/terminalMaze/logger"
)

type Player struct {
	name              string
	initHealthPoint   int
	initRationalPoint int
	initAttackPoint   int
	healthPoint       int
	rationalPoint     int
	attackPoint		  int	
	itemsList         map[string]int
	preloads		  []PlayerMiddleware
	callbacks 		  []PlayerMiddleware
}

type PlayerConfig struct {
	Name              string
	InitHealthPoint   int
	InitRationalPoint int
	InitAttackPoint   int
}

func NewPlayer(config *PlayerConfig) *Player {
	return &Player{
		name: config.Name,
		initHealthPoint: config.InitHealthPoint,
		initRationalPoint: config.InitRationalPoint,
		initAttackPoint: config.InitAttackPoint,
		healthPoint: config.InitHealthPoint,
		rationalPoint: config.InitRationalPoint,
		attackPoint: config.InitAttackPoint,
		itemsList: make(map[string]int),
		preloads: make([]PlayerMiddleware, 0),
		callbacks: make([]PlayerMiddleware, 0),
	}
}

func (player *Player) GetClass() reflect.Type {
	return reflect.TypeOf(player)
}

func (player *Player) ToString() string {
	ans := fmt.Sprintf("%v", player)
	return ans
}

func (player *Player) Copy() *Player {
	return nil
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) GetHealthPoint() int {
	return player.healthPoint
}

func (player *Player) GetRationalPoint() int {
	return player.rationalPoint
}

func (player *Player) GetAttackPoint() int {
	return player.attackPoint
}

func (player *Player) DisplayItems(logger *logger.Logger) {
	for key, value := range player.itemsList {
		logger.Println("[Item]<" + key + "> x " + strconv.Itoa(value))
	}
}

func (player *Player) AddItems(itemName string, num int) *Player {
	if num < 0 {
		return player
	}

	player.itemsList[itemName] += num
	return player
}

func (player *Player) UseItem(itemName string) string {
	if _, ok := player.itemsList[itemName]; !ok {
		return ""
	}

	if player.itemsList[itemName] == 0 {
		return ""
	}

	player.itemsList[itemName]--
	if player.itemsList[itemName] == 0 {
		delete(player.itemsList, itemName)
	}

	return itemName
}

func (player *Player) GetInitHealthPoint() int {
	return player.initHealthPoint
}

func (player *Player) GetInitRationalPoint() int {
	return player.initRationalPoint
}

func (player *Player) GetInitAttackPoint() int {
	return player.initAttackPoint
}

func (player *Player) UpdateHeathPoint(change int) (*Player, bool) {
	currentHealthPoint := player.GetHealthPoint()
	if change + currentHealthPoint <= 0 {
		player.healthPoint = 0
		return player, true
	}

	player.healthPoint += change
	return player, false
}

func (player *Player) UpdateRationalPoint(change int) (*Player, bool) {
	currentRationalPoint := player.GetRationalPoint()
	if change + currentRationalPoint <= 0 {
		player.rationalPoint = 0
		return player, true
	}

	player.rationalPoint += change
	return player, false
}

func (player *Player) UpdateAttackPoint(change int) *Player {
	currentAttackPoint := player.GetAttackPoint()
	if change + currentAttackPoint <= 0 {
		player.attackPoint = 0
		return player
	}

	player.attackPoint += change
	return player
}

func (player *Player) AppendPreloads(middlewares ...PlayerMiddleware) *Player {
	player.preloads = append(player.preloads, middlewares...)
	return player
}

func (player *Player) AppendCallbacks(middlewares ...PlayerMiddleware) *Player {
	player.callbacks = append(player.callbacks, middlewares...)
	return player
}

func (player *Player) ClearPreloads() *Player {
	player.preloads = make([]PlayerMiddleware, 0)
	return player
}

func (player *Player) ClearCallbacks() *Player {
	player.callbacks = make([]PlayerMiddleware, 0)
	return player
}

func (player *Player) ApplyPreloads(logger *logger.Logger) {
	if logger.IsDead() || logger.IsOver() {
		return
	}

	for _, handler := range player.preloads {
		handler(player, logger)
	}
}

func (player *Player) ApplyCallbacks(logger *logger.Logger) {
	if logger.IsDead() || logger.IsOver() {
		return
	}

	for _, handler := range player.callbacks {
		handler(player, logger)
	}
}




