package concreteHandlerB

// handler provides a handler interface
type handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerB provides a ConcreteHandlerB interface
type ConcreteHandlerB interface {
	SendRequest(int) string
}

// concreteHandlerB implements concreteHandlerB
type concreteHandlerB struct {
	next handler
}

// SendRequest implementation
func (h *concreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler B"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// NewConcreteHandlerB creates ConcreteHandlerB
func NewConcreteHandlerB(h handler) ConcreteHandlerB {
	return &concreteHandlerB{next: h}
}
