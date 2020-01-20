package visitor

import "fmt"

// RedSquare provides the RedSquare interface
type RedSquare interface {
	Accept(visitor) string
	MakePhoto() string
}

// redSquare implements the place interface.
type redSquare struct {
	name      string
	showplace string
}

// Accept implementation.
func (r *redSquare) Accept(v visitor) string {
	return v.VisitRedSuare(r)
}

// MakePhoto implementation
func (r *redSquare) MakePhoto() string {
	return fmt.Sprintf("Make a photo of %v in %v\n", r.showplace, r.name)
}

// NewRedSquare creates redSquare
func NewRedSquare() RedSquare {
	return &redSquare{
		name:      "The Red Square",
		showplace: "Spasskaya Tower",
	}
}
