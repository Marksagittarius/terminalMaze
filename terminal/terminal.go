package terminal

type TerminalController interface {
	ReadDirection() int
	ReadMenuCommand() (string, string)
}

