package visitor

import "fmt"

// bolshoiTheatre implements the place interface.
type bolshoiTheatre struct {
	name    string
	product string
}

// Accept implementation.
func (b *bolshoiTheatre) Accept(v visitor) string {
	return v.VisitBolhoiTheatre(b)
}

// SeePerformance implementation.
func (b *bolshoiTheatre) SeePerformance() string {
	return fmt.Sprintf("See a %v in %v\n", b.product, b.name)
}

func NewBolshoiTheatre() *bolshoiTheatre {
	return &bolshoiTheatre{
		name:    "The Bolshoi Theatre",
		product: "performance",
	}
}
