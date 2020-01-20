package visitor

import "fmt"

// GorkyPark provides the GorkyPark interface
type GorkyPark interface {
	Accept(visitor) string
	WalkAround() string
}

// gorkyPark implements the place interface.
type gorkyPark struct {
	name    string
	product string
}

// Accept implementation.
func (g *gorkyPark) Accept(v visitor) string {
	return v.VisitGorkyPark(g)
}

// WalkAround implementation.
func (g *gorkyPark) WalkAround() string {
	return fmt.Sprintf("Walk among %v in %v\n", g.product, g.name)
}

// NewGorkyPark creates gorkyPark
func NewGorkyPark() GorkyPark {
	return &gorkyPark{
		name:    "The Gorky Park",
		product: "trees and fountains",
	}
}
