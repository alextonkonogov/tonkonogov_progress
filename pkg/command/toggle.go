package command

import "fmt"

// Toggle provides a Toggle interface
type Toggle interface {
	StoreCommand(executor)
	ExecuteCommands() string
}

// toggle implements a toggle (invoker)
type toggle struct {
	commands []executor
}

// StoreCommand implementation
func (t *toggle) StoreCommand(cmd executor) {
	t.commands = append(t.commands, cmd)
}

// ExecuteCommands implementation
func (t *toggle) ExecuteCommands() (s string) {
	for _, v := range t.commands {
		s += fmt.Sprintf("%v\n", v.Execute())
	}
	return
}

// NewToggle creates a toggle
func NewToggle() Toggle {
	return &toggle{}
}
