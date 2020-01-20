package switchOnCommand

// lamp provides a lamp interface
type lamp interface {
	TurnOn() string
}

// SwitchOnCommand provides a SwitchOnCommand interface
type SwitchOnCommand interface {
	Execute() string
}

// switchOnCommand implements switchOnCommand - concrete command for lamp to turn on the light
type switchOnCommand struct {
	lamp lamp
}

// execute implementation
func (c *switchOnCommand) Execute() string {
	return c.lamp.TurnOn()
}

// NewSwitchOnCommand creates switchOnCommand
func NewSwitchOnCommand(lamp lamp) SwitchOnCommand {
	return &switchOnCommand{lamp: lamp}
}
