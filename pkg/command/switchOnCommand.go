package command

// switchOnCommand implements switchOnCommand - concrete command for lamp to turn on the light
type switchOnCommand struct {
	lamp *lamp
}

// execute implementation
func (c *switchOnCommand) execute() string {
	return c.lamp.TurnOn()
}

// NewSwitchOnCommand creates switchOnCommand
func NewSwitchOnCommand(lamp *lamp) *switchOnCommand {
	return &switchOnCommand{lamp: lamp}
}
