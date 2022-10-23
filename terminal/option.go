package terminal

import (
	"fmt"
	"strings"
)

type OptionElement struct {
	Key     string
	Content string
}

type OptionSelector struct {
	options map[string]*OptionElement
}

func NewOptionSelector() *OptionSelector {
	return &OptionSelector{
		options: make(map[string]*OptionElement, 0),
	}
}

func (os *OptionSelector) AppendOption(options ...*OptionElement) *OptionSelector {
	for _, option := range options {
		os.options[option.Key] = option
	}

	return os
}

func (os *OptionSelector) DisplayOptions() {
	for key, value := range os.options {
		fmt.Println("(" + key + ") " + value.Content)
	}
}

func (os *OptionSelector) ReadSelection() (string, string) {
	if len(os.options) == 0 {
		return "", ""
	}

	commandTimes := 0
	selectResult := ""
	fmt.Print(">> ")
	option, ok := os.options[strings.ToLower(selectResult)]
	for !ok {
		if commandTimes > 1 {
			fmt.Println("[Warning] Wrong parameter in command <select>.")
			fmt.Print(">> ")
		}
		commandTimes++
		fmt.Scanf(ReadDirectionCommand, &selectResult)
		option, ok = os.options[strings.ToLower(selectResult)]
	}

	return option.Key, option.Content
}