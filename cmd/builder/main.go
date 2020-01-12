package main

import (
	"fmt"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/builder/cleanUpTheRoom"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/builder/mother"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/builder/order"
)

var toys, bed, floor = "scattered", "prorerly", "normally"

func main() {
	order := order.NewOrder()
	mother := mother.NewMother(cleanUpTheRoom.NewRoomCleaning(order))
	mother.Cleaning(toys, bed, floor)
	result := order.Show()

	fmt.Println(result)
}
