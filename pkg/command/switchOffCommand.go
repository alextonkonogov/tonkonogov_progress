package command

// SwitchOffCommand - ConcreteCommand(The executor for turning off the light)
type switchOffCommand struct {
	lamp *lamp
}

// execute implementation
func (c *switchOffCommand) execute() string {
	return c.lamp.TurnOff()
}

// NewSwitchOnCommand creates switchOnCommand
func NewSwitchOffCommand(lamp *lamp) *switchOffCommand {
	return &switchOffCommand{lamp: lamp}
}
