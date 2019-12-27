package main

import (
	"fmt"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/builder"
)

var toys, bed, floor = "scattered", "prorerly", "normally"

func main() {
	order := builder.NewOrder()
	builder.CallThis(order)
	mother := builder.NewMother(builder.NewRoomCleaning(order))
	builder.CallThis(mother)
	builder.CallThis(builder.NewRoomCleaning(order))
	mother.Cleaning(toys, bed, floor)
	result := order.Show()

	fmt.Println(result)
}
