package builder

import "fmt"

// son provides a son interface
type son interface {
	CollectToys(direction string)
	MakeTheBed(direction string)
	VacuumTheFloor(direction string)
}

// mother implements a manager
type mother struct {
	son son
}

// NewMother creates mother
func NewMother(son son) *mother {
	return &mother{son: son}
}

// Cleaning tells the son what to do, in what order and how
func (m *mother) Cleaning(toys, bed, floor string) {
	m.son.CollectToys(toys)
	m.son.MakeTheBed(bed)
	m.son.VacuumTheFloor(floor)
}

// cleanUpTheRoom implements son interface
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

// order implements an order
type order struct {
	result string
}

// Show demonstrates an order in the room
func (o *order) Show() string {
	return o.result
}

// NewOrder creates an order
func NewOrder() *order {
	return &order{}
}
