package lamp

// Lamp provides a Lamp interface
type Lamp interface {
	TurnOn() string
	TurnOff() string
}

// light implements a light (receiver)
type lamp struct{}

// TurnOn implementation
func (l *lamp) TurnOn() string {
	return "The light is on"
}

// TurnOff implementation
func (l *lamp) TurnOff() string {
	return "The light is off"
}

// NewLamp creates lamp
func NewLamp() Lamp {
	return &lamp{}
}
