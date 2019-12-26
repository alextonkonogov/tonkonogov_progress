package main

import (
	"fmt"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/visitor"
)

func main() {
	city := visitor.NewCity()

	city.Add(visitor.NewRedSquare())
	city.Add(visitor.NewBolshoiTheatre())
	city.Add(visitor.NewBurgerBar())

	result := city.Accept(visitor.NewTourist())
	fmt.Println(result)
}
