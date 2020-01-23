package main

import (
	"fmt"
	cr "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/creater"
)

func main() {

	factory := cr.NewCreater()
	products := []cr.Producter{
		factory.CreateProduct("A"),
		factory.CreateProduct("B"),
		factory.CreateProduct("C"),
	}

	for _, product := range products {
		fmt.Println(product.Use())
	}
}
