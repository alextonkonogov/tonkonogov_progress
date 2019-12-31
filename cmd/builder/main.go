package main

import (
	"fmt"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/builder"
)

var toys, bed, floor = "scattered", "prorerly", "normally"

func main() {
	order := builder.NewOrder()
	builder.Call(order)
	mother := builder.NewMother(builder.NewRoomCleaning(order))
	builder.Call(mother)
	builder.Call(builder.NewRoomCleaning(order))
	mother.Cleaning(toys, bed, floor)
	result := order.Show()

	fmt.Println(result)
}
