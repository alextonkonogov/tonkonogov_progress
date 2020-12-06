package whip

// Whip provides a Whip interface
type Whip interface {
	Use() string
}

// whip implements a whip interface
type whip struct {
	action string
}

// Use implementation
func (w *whip) Use() string {
	return w.action
}

// NewWhip creates a Whip
func NewWhip() Whip {
	return &whip{
		action: "whoosh whoosh!",
	}
}
