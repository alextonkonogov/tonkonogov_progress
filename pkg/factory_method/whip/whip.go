package whip

type Whip interface {
	Use() string
}

type whip struct {
	action string
}

func (w *whip) Use() string {
	return w.action
}

func NewWhip() Whip {
	return &whip{
		action: "whoosh whoosh!",
	}
}
