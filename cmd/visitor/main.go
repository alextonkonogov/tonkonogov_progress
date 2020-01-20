package main

import (
	"fmt"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/visitor"
)

func main() {
	city := visitor.NewMoscow()
	city.Add(visitor.NewRedSquare())
	city.Add(visitor.NewBolshoiTheatre())
	city.Add(visitor.NewGorkyPark())
	result := city.Accept(visitor.NewTourist())
	fmt.Println(result)
}
