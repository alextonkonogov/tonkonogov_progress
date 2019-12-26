package visitor

// City implements a collection of places to visit.
type moscow struct {
	places []place
}

// Add appends place to the collection.
func (m *moscow) Add(p place) {
	m.places = append(m.places, p)
}

// Accept implements a visit to all places in the moscow.
func (m *moscow) Accept(v visitor) string {
	var result string
	for _, p := range m.places {
		result += p.Accept(v)
	}
	return result
}

func NewCity() *moscow {
	return &moscow{
		places: []place{},
	}
}
