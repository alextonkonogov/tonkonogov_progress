package visitor

// visitor provides a visitor interface.
type visitor interface {
	VisitRedSuare(p *redSquare) string
	VisitBolhoiTheatre(p *bolshoiTheatre) string
	VisitGorkyPark(p *gorkyPark) string
}

// place provides an interface for place that the visitor should visit.
type place interface {
	Accept(v visitor) string
}
