package concreteHandlerA

// handler provides a handler interface
type handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerA provides a ConcreteHandlerA interface
type ConcreteHandlerA interface {
	SendRequest(int) string
}

// concreteHandlerA implements concreteHandlerA
type concreteHandlerA struct {
	next handler
}

// SendRequest implementation
func (h *concreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler A"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// NewConcreteHandlerA creates ConcreteHandlerA
func NewConcreteHandlerA(h handler) ConcreteHandlerA {
	return &concreteHandlerA{next: h}
}
