package sword

type Sword interface {
	Use() string
}

// Sword implements product "A"
type sword struct {
	action string
}

// Use returns product action
func (s *sword) Use() string {
	return s.action
}

func NewSword() Sword {
	return &sword{
		action: "shuh shuh shuh!",
	}
}
