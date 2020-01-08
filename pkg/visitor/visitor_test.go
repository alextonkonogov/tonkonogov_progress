package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	expect := "Make a photo of Spasskaya Tower in The Red Square\nSee a performance in The Bolshoi Theatre\nWalk among trees and fountains in The Gorky Park\n"

	city := NewMoscow()
	city.Add(NewRedSquare())
	city.Add(NewBolshoiTheatre())
	city.Add(NewGorkyPark())

	result := city.Accept(NewTourist())

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
