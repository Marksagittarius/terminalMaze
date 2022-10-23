package terminal

import (
	"fmt"
	"strings"
	"github.com/Marksagittarius/terminalMaze/logger"
)

type DefaultTerminalController struct {
	logger *logger.Logger
}

func NewDefaultTerminalController(logger *logger.Logger) *DefaultTerminalController {
	return &DefaultTerminalController{
		logger: logger,
	}
}

func (dtc *DefaultTerminalController) ReadDirection() int {
	commandTimes := 0
	direction := ""
	fmt.Print(">> ")
	directionCode, ok := DirectionDictionary[strings.ToLower(direction)]
	for !ok {
		if commandTimes > 1 {
			dtc.logger.Println("[Warning] Wrong parameter in command <move>.")
			fmt.Print(">> ")
		}
		commandTimes++
		fmt.Scanf(ReadDirectionCommand, &direction)
		directionCode, ok = DirectionDictionary[strings.ToLower(direction)]
	}

	return directionCode
}

func (dtc *DefaultTerminalController) ReadMenuCommand() (string, string) {
	commandTimes := 0
	commandLeft := ""
	commandRight := ""
	fmt.Print(">> ")

	for  {
		if commandTimes > 1 {
			dtc.logger.Println("[Warning] Wrong parameter in menu command.")
			fmt.Print(">> ")
		}
		commandTimes++
		fmt.Scanf("%s", &commandLeft)

		if commandLeft == "display" {
			fmt.Scanf("%s", &commandRight)
			if commandRight == "items" {
				return "display", "items"
			}
		} else if commandLeft == "use" {
			fmt.Scanf("%s", &commandRight)
			return commandLeft, commandRight
		} else if commandLeft == "move" {
			return "move", ""
		}
	}
}

