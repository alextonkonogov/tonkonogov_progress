package switchOffComand

// lamp provides a lamp interface
type lamp interface {
	TurnOff() string
}

// SwitchOffCommand provides a SwitchOffCommand interface
type SwitchOffCommand interface {
	Execute() string
}

// SwitchOffCommand - ConcreteCommand(The executor for turning off the light)
type switchOffCommand struct {
	lamp lamp
}

// execute implementation
func (c *switchOffCommand) Execute() string {
	return c.lamp.TurnOff()
}

// NewSwitchOnCommand creates switchOnCommand
func NewSwitchOffCommand(lamp lamp) SwitchOffCommand {
	return &switchOffCommand{lamp: lamp}
}
