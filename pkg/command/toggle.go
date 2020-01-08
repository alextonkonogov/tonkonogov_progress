package command

import "fmt"

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
		s += fmt.Sprintf("%v\n", v.execute())
	}
	return
}

// NewToggle creates toggle
func NewToggle() *toggle {
	return &toggle{}
}
