package builder

import "fmt"

// roomCleaning implements child interface
type roomCleaning struct {
	order *order
}

// MakeTheBed makes the bed
func (r *roomCleaning) MakeTheBed(direction string) {
	r.order.result += fmt.Sprintf("Made the bed %v\n", direction)
}

// CollectToys collects the toys
func (r *roomCleaning) CollectToys(direction string) {
	r.order.result += fmt.Sprintf("Collected %v toys\n", direction)
}

// VacuumTheFloor cleans the floor using vacuum cleaner
func (r *roomCleaning) VacuumTheFloor(direction string) {
	r.order.result += fmt.Sprintf("Vacuumed the floor %v\n", direction)
}

func (r *roomCleaning) Call() {
	fmt.Println("...starting of room cleaning")
}

// NewRoomCleaning creates roomCleaning
func NewRoomCleaning(order *order) *roomCleaning {
	return &roomCleaning{
		order: order,
	}
}
