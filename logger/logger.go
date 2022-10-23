package logger

import (
	"fmt"
	"github.com/Marksagittarius/terminalMaze/redux"
)

type Logger struct {
	trace      []string
	log        []string
	operations []int
	isDead     bool
	isOver     bool
	tokenList  *redux.Store
}

func NewLogger() *Logger {
	defaultReducer := func(state *redux.State, action redux.Action) *redux.State {
		return action.Payload(state)
	}

	return &Logger{
		trace:      make([]string, 0),
		log:        make([]string, 0),
		isDead:     false,
		isOver:     false,
		tokenList:  redux.CreateStore(defaultReducer, &redux.State{
			HashMap: make(map[string]any),
		}),
		operations: make([]int, 0),
	}
}

func (logger *Logger) IsDead() bool {
	return logger.isDead
}

func (logger *Logger) IsOver() bool {
	return logger.isOver
}

func (logger *Logger) CondemnToDeath() *Logger {
	logger.isDead = true
	return logger
}

func (logger *Logger) GameOver() *Logger {
	logger.isOver = true
	return logger
}

func (logger *Logger) Println(info string) *Logger {
	logger.log = append(logger.log, info)
	fmt.Println(info)
	return logger
}

func (logger *Logger) LoadPrinter(middleware LoggerPrinterMiddleware, input string) *Logger {
	output := middleware(logger, input)
	logger.log = append(logger.log, output)
	return logger
}

func (logger *Logger) AppendOperation(operation int) *Logger {
	logger.operations = append(logger.operations, operation)
	return logger
} 

func (logger *Logger) GetLastOperation() int {
	if len(logger.operations) == 0 {
		return -1
	}

	return logger.operations[len(logger.operations) - 1]
}

func (logger *Logger) GetState(tokenName string) string {
	state, err := logger.tokenList.GetState()
	if err != nil {
		return ""
	}

	if _, ok := state.HashMap[tokenName]; !ok {
		return ""
	}

	ans := fmt.Sprintf("%v", state.HashMap[tokenName])
	return ans
}

func (logger *Logger) Subscribe(listeners ...redux.Listener) (func() error, error) {
	return logger.tokenList.Subscribe(listeners...)
}

func (logger *Logger) Dispatch(actions ...redux.Action) ([] redux.Action, error) {
	return logger.tokenList.Dispatch(actions...)
}

func (logger *Logger) ReplaceReducer(nextReducer redux.Reducer) error {
	return logger.tokenList.ReplaceReducer(nextReducer)
}

func (logger *Logger) ApplyMiddlewares(middlewares... redux.Middleware) (func() error, error) {
	return logger.tokenList.ApplyMiddlewares(middlewares...)
}




