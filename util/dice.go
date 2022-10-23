package util

import (
	"math/rand"
	"time"
)

type Dice struct {
	maxValue int
}

func NewDice(max int) *Dice {
	return &Dice{
		maxValue: max,
	}
}

func (dice *Dice) Roll() int {
	rand.Seed(time.Now().Unix())
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(dice.maxValue) + 1
}