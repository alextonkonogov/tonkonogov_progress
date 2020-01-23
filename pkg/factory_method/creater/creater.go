package creater

// Producter provides a product interface.
// All products returned by factory must provide a single interface.
type Producter interface {
	Use() string // Every product should be can be used
}

// Creater provides a factory interface.
type Creater interface {
	CreateProduct(producter Producter) Producter // Factory Method
}

// ConcreteCreater implements Creater interface.
type creater struct{}

// CreateProduct is a Factory Method
func (p *creater) CreateProduct(producter Producter) Producter {
	return producter
}

// NewCreater is the ConcreteCreater constructor.
func NewCreater() Creater {
	return &creater{}
}
