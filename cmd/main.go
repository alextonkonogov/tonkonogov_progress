package main

import (
	"fmt"

	cr "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/creater"
)

func main() {
	factory := cr.NewBlacksmith()
	products := []cr.Weapon{
		factory.CreateWeapon("sword"),
		factory.CreateWeapon("hammer"),
		factory.CreateWeapon("whip"),
	}

	for _, product := range products {
		fmt.Println(product.Use())
	}
}
