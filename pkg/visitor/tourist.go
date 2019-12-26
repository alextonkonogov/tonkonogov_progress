package visitor

// tourist implements the visitor interface.
type tourist struct {
}

// VisitBolhoiTheatre implements visit to SushiBar.
func (v *tourist) VisitBolhoiTheatre(p *bolshoiTheatre) string {
	return p.SeePerformance()
}

// VisitRedSuare implements visit to Pizzeria.
func (v *tourist) VisitRedSuare(p *redSquare) string {
	return p.MakePhoto()
}

// VisitGorkyPark implements visit to BurgerBar.
func (v *tourist) VisitGorkyPark(p *gorkyPark) string {
	return p.WalkAround()
}

func NewTourist() *tourist {
	return &tourist{}
}
