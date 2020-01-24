package concreteHandlerC

// handler provides a handler interface
type handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerС provides a ConcreteHandlerС interface
type ConcreteHandlerC interface {
	SendRequest(int) string
}

// concreteHandlerC implements concreteHandlerC
type concreteHandlerC struct {
	next handler
}

// SendRequest implementation
func (h *concreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler C"
	} else {
		result = "...silence..."
	}
	return
}

// NewConcreteHandlerC creates ConcreteHandlerC
func NewConcreteHandlerC() ConcreteHandlerC {
	return &concreteHandlerC{}
}
