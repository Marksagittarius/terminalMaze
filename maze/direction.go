package maze

type Direction int

const (
	Null  Direction = -1
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

func GetReverseDirection(direction Direction) Direction {
	switch direction {
		case North:
			return South
		case South:
			return North
		case East:
			return West
		case West:
			return East
	}

	return Null
}