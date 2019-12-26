package visitor

// tourist implements the visitor interface.
type tourist struct {
}

// VisitBolhoiTheatre implements visit to BolhoiTheatre
func (v *tourist) VisitBolhoiTheatre(p *bolshoiTheatre) string {
	return p.SeePerformance()
}

// VisitRedSuare implements visit to RedSuare
func (v *tourist) VisitRedSuare(p *redSquare) string {
	return p.MakePhoto()
}

// VisitGorkyPark implements visit to GorkyPark
func (v *tourist) VisitGorkyPark(p *gorkyPark) string {
	return p.WalkAround()
}

func NewTourist() *tourist {
	return &tourist{}
}
