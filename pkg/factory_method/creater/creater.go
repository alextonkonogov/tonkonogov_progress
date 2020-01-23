package creater

import (
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/productA"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/productB"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/productC"
	"log"
)

// Producter provides a product interface.
// All products returned by factory must provide a single interface.
type Producter interface {
	Use() string // Every product should be can be used
}

// Creater provides a factory interface.
type Creater interface {
	CreateProduct(string) Producter // Factory Method
}

// ConcreteCreater implements Creater interface.
type creater struct{}

// CreateProduct is a Factory Method
func (p *creater) CreateProduct(action string) Producter {
	var product Producter

	switch action {
	case "A":
		product = productA.NewConcreteProductA(action)
	case "B":
		product = productB.NewConcreteProductB(action)
	case "C":
		product = productC.NewConcreteProductC(action)
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

// NewCreater is the ConcreteCreater constructor.
func NewCreater() Creater {
	return &creater{}
}
