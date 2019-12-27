package builder

import "fmt"

// order implements an order
type order struct {
	result string
}

// Show demonstrates an order in the room
func (o *order) Show() string {
	return o.result
}

func (o *order) Call() {
	fmt.Println("...order is expected")
}

// NewOrder creates an order
func NewOrder() *order {
	return &order{}
}
