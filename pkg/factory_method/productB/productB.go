package productB

type ConcreteProductB interface {
	Use() string
}

type concreteProductB struct {
	action string
}

func (p *concreteProductB) Use() string {
	return p.action
}

func NewConcreteProductB(action string) ConcreteProductB {
	return &concreteProductB{
		action: action,
	}
}
