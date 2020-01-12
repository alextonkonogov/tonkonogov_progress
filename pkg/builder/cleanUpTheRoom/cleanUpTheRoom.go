package cleanUpTheRoom

import (
	"fmt"
)

// roomCleaning implements child interface
type roomCleaning struct {
	order order
}

// MakeTheBed makes the bed
func (r *roomCleaning) MakeTheBed(direction string) {
	result := fmt.Sprintf("Made the bed %v\n", direction)
	r.order.AddResult(result)
}

// CollectToys collects the toys
func (r *roomCleaning) CollectToys(direction string) {
	result := fmt.Sprintf("Collected %v toys\n", direction)
	r.order.AddResult(result)
}

// VacuumTheFloor cleans the floor using vacuum cleaner
func (r *roomCleaning) VacuumTheFloor(direction string) {
	result := fmt.Sprintf("Vacuumed the floor %v\n", direction)
	r.order.AddResult(result)
}

// order provides an order interface
type order interface {
	AddResult(string)
}

// NewRoomCleaning creates roomCleaning
func NewRoomCleaning(order order) *roomCleaning {
	return &roomCleaning{
		order: order,
	}
}
