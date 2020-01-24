package hammer

type Hammer interface {
	Use() string
}

type hammer struct {
	action string
}

func (h *hammer) Use() string {
	return h.action
}

func NewHammer() Hammer {
	return &hammer{
		action: "bang!",
	}
}
