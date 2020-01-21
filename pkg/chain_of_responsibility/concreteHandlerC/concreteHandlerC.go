package concreteHandlerC

// handler provides a handler interface
type handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerC implements concreteHandlerC.
type ConcreteHandlerC interface {
	SendRequest(int) string
}

type concreteHandlerC struct {
	next handler
}

// SendRequest implementation.
func (h *concreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler C"
	} else {
		result = "...silence..."
	}
	return
}

func NewConcreteHandlerC() ConcreteHandlerC {
	return &concreteHandlerC{}
}
