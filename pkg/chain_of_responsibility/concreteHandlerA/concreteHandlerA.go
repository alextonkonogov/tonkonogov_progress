package concreteHandlerA

// handler provides a handler interface
type handler interface {
	SendRequest(message int) string
}

// concreteHandlerA implements handler concreteHandlerA
type ConcreteHandlerA interface {
	SendRequest(int) string
}

type concreteHandlerA struct {
	next handler
}

// SendRequest implementation.
func (h *concreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler A"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

func NewConcreteHandlerA(h handler) ConcreteHandlerA {
	return &concreteHandlerA{next: h}
}
