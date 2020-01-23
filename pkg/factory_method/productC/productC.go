package productC

type ConcreteProductC interface {
	Use() string
}

type concreteProductC struct {
	action string
}

func (p *concreteProductC) Use() string {
	return p.action
}

func NewConcreteProductC(action string) ConcreteProductC {
	return &concreteProductC{
		action: action,
	}
}
