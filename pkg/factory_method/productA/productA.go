package productA

type ConcreteProductA interface {
	Use() string
}

// ConcreteProductA implements product "A"
type concreteProductA struct {
	action string
}

// Use returns product action
func (p *concreteProductA) Use() string {
	return p.action
}

func NewConcreteProductA(action string) ConcreteProductA {
	return &concreteProductA{
		action: action,
	}
}
