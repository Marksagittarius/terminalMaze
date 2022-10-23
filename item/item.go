package item

import (
	"fmt"
	"reflect"
	"github.com/Marksagittarius/terminalMaze/factory"
	"github.com/Marksagittarius/terminalMaze/logger"
	"github.com/Marksagittarius/terminalMaze/player"
)

type Item struct {
	itemName      string
	rationalPoint int
	effect        ItemMiddleware
}

func (item *Item) GetItemName() string {
	return item.itemName
}

func (item *Item) GetRationalPoint() int {
	return item.rationalPoint
}

func (item *Item) SetEffect(effect ItemMiddleware) ItemPrototype {
	item.effect = effect
	return item
}

func (item *Item) GetClass() reflect.Type {
	return reflect.TypeOf(item)
}

func (item *Item) ToString() string {
	ans := fmt.Sprintf("%v", item)
	return ans
}

func (item *Item) Copy() factory.AbstractProduct {
	copiedItemName := item.itemName
	copiedRationalPoint := item.rationalPoint
	copiedEffect := item.effect

	return &Item{
		itemName: copiedItemName,
		rationalPoint: copiedRationalPoint,
		effect: copiedEffect,
	}
}

func (item *Item) UseEffect(player *player.Player, logger *logger.Logger) {
	if item.effect == nil {
		return
	}

	item.effect(player, logger)
}
