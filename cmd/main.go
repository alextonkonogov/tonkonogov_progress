package main

import (
	"fmt"
	cr "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/creater"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/productA"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/productB"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/productC"
	"log"
)

func main() {

	var f = func(action string) cr.Producter {
		var product cr.Producter

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

	factory := cr.NewCreater()
	products := []cr.Producter{
		factory.CreateProduct(f("A")),
		factory.CreateProduct(f("B")),
		factory.CreateProduct(f("C")),
	}

	for _, product := range products {
		fmt.Println(product.Use())
	}
}
