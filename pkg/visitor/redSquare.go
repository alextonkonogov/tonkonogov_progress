package visitor

import "fmt"

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

func NewRedSquare() *redSquare {
	return &redSquare{
		name:      "The Red Square",
		showplace: "Spasskaya Tower",
	}
}
