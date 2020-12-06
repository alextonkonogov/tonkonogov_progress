package sword

// Sword provides a Sword interface
type Sword interface {
	Use() string
}

// sword implements a sword interface
type sword struct {
	action string
}

// Use implementation
func (s *sword) Use() string {
	return s.action
}

// NewSword creates a Sword
func NewSword() Sword {
	return &sword{
		action: "shuh shuh shuh!",
	}
}
