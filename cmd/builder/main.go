package main

import (
	"fmt"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/builder"
)

var toys, bed, floor = "scattered", "prorerly", "normally"

func main() {
	order := builder.NewOrder()
	mother := builder.NewMother(builder.NewCleanUpTheRoom(order))
	mother.Cleaning(toys, bed, floor)
	result := order.Show()

	fmt.Println(result)
}
