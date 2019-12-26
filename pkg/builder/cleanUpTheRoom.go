package builder

import "fmt"

// cleanUpTheRoom implements child interface
type cleanUpTheRoom struct {
	order *order
}

// NewCleanUpTheRoom creates cleanUpTheRoom
func NewCleanUpTheRoom(order *order) *cleanUpTheRoom {
	return &cleanUpTheRoom{
		order: order,
	}
}

// MakeTheBed makes the bed
func (c *cleanUpTheRoom) MakeTheBed(direction string) {
	c.order.result += fmt.Sprintf("Made the bed %v\n", direction)
}

// CollectToys collects the toys
func (c *cleanUpTheRoom) CollectToys(direction string) {
	c.order.result += fmt.Sprintf("Collected %v toys\n", direction)
}

// VacuumTheFloor cleans the floor using vacuum cleaner
func (c *cleanUpTheRoom) VacuumTheFloor(direction string) {
	c.order.result += fmt.Sprintf("Vacuumed the floor %v\n", direction)
}
