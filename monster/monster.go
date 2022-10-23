package monster

import (
	"fmt"
	"reflect"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/player"
)

type Monster struct {
	monsterName     string
	initHealthPoint int
	initAttackPoint int
	healthPoint     int
	attackPoint     int
	preloads        []MonsterMiddleware
	callbacks		[]MonsterMiddleware
}

func (monster *Monster) GetMonsterName() string {
	return monster.monsterName
}

func (monster *Monster) GetClass() reflect.Type {
	return reflect.TypeOf(monster)
}

func (monster *Monster) ToString() string {
	ans := fmt.Sprintf("%v", monster)
	return ans
}

func (monster *Monster) Copy() MonsterPrototype {
	copiedMonsterName := monster.monsterName
	copiedInitHealthPoint := monster.initHealthPoint
	copiedInitAttackPoint := monster.initAttackPoint
	copiedHealthPoint := monster.healthPoint
	copiedAttackPoint := monster.attackPoint

	copiedPreloads := make([]MonsterMiddleware, len(monster.preloads))
	copy(copiedPreloads, monster.preloads)
	copiedCallbacks := make([]MonsterMiddleware, len(monster.callbacks))
	copy(copiedCallbacks, monster.callbacks)

	return &Monster{
		monsterName: copiedMonsterName,
		initHealthPoint: copiedInitHealthPoint,
		initAttackPoint: copiedInitAttackPoint,
		healthPoint: copiedHealthPoint,
		attackPoint: copiedAttackPoint,
		preloads: copiedPreloads,
		callbacks: copiedCallbacks,
	}
}

func (monster *Monster) GetHealthPoint() int {
	return monster.healthPoint
}

func (monster *Monster) GetAttackPoint() int {
	return monster.attackPoint
}

func (monster *Monster) GetInitHealthPoint() int {
	return monster.initHealthPoint
}

func (monster *Monster) GetInitAttackPoint() int {
	return monster.initAttackPoint
}

func (monster *Monster) AppendPreloads(middlewares... MonsterMiddleware) MonsterPrototype {
	monster.preloads = append(monster.preloads, middlewares...)
	return monster
}

func (monster *Monster) AppendCallbacks(middlewares... MonsterMiddleware) MonsterPrototype {
	monster.callbacks = append(monster.callbacks, middlewares...)
	return monster
}

func (monster *Monster) ClearPreloads() MonsterPrototype {
	monster.preloads = make([]MonsterMiddleware, 0)
	return monster
}

func (monster *Monster) ClearCallbacks() MonsterPrototype {
	monster.callbacks = make([]MonsterMiddleware, 0)
	return monster
}

func (monster *Monster) UpdateHeathPoint(change int) (MonsterPrototype, bool) {
	currentHealthPoint := monster.GetHealthPoint()
	if currentHealthPoint + change <= 0 {
		monster.healthPoint = 0
		return monster, true
	}

	monster.healthPoint += change
	return monster, false
}

func (monster *Monster) UpdateAttackPoint(change int) MonsterPrototype {
	currentAttackPoint := monster.GetAttackPoint()
	if currentAttackPoint + change <= 0 {
		monster.attackPoint = 0
		return monster
	}

	monster.attackPoint += change
	return monster
}

func (monster *Monster) ApplyPreloads(player *player.Player, logger *logger.Logger) {
	if logger.IsDead() || logger.IsOver() {
		return
	}
	
	for _, handler := range monster.preloads {
		handler(monster, player, logger)
	}
}

func (monster *Monster) ApplyCallbacks(player *player.Player, logger *logger.Logger) {
	if logger.IsDead() || logger.IsOver() {
		return
	}
	
	for _, handler := range monster.preloads {
		handler(monster, player, logger)
	}
}






