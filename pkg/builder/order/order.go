package order

// Order implements an Order public interface
type Order interface {
	AddResult(string)
	Show() string
}

// order implements an order
type order struct {
	result string
}

// AddResult adds
func (o *order) AddResult(result string) {
	o.result += result
}

// Show demonstrates an order in the room
func (o *order) Show() string {
	return o.result
}

// NewOrder creates an order
func NewOrder() Order {
	return &order{}
}
