package hammer

// Hammer provides a Hammer interface
type Hammer interface {
	Use() string
}

// hammer implements a hammer interface
type hammer struct {
	action string
}

// Use implementation
func (h *hammer) Use() string {
	return h.action
}

// NewHammer creates a Hammer
func NewHammer() Hammer {
	return &hammer{
		action: "bang!",
	}
}
