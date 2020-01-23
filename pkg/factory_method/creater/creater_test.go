package creater

import (
	"log"
	"testing"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/productA"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/productB"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/productC"
)

func TestFactoryMethod(t *testing.T) {

	assert := []string{"A", "B", "C"}

	var f = func(action string) producter {
		var product producter

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

	factory := NewCreater()
	products := []producter{
		factory.CreateProduct(f("A")),
		factory.CreateProduct(f("B")),
		factory.CreateProduct(f("C")),
	}

	for i, product := range products {
		if action := product.Use(); action != assert[i] {
			t.Errorf("Expect action to %s, but %s.\n", assert[i], action)
		}
	}
}
