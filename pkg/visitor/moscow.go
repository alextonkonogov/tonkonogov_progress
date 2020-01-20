package visitor

// Moscow provides Moscow interface
type Moscow interface {
	Add(place)
	Accept(visitor) string
}

// moscow implements a collection of places to visit.
type moscow struct {
	places []place
}

// Add appends place to the collection.
func (m *moscow) Add(p place) {
	m.places = append(m.places, p)
}

// Accept implements a visit to all places in moscow.
func (m *moscow) Accept(v visitor) string {
	var result string
	for _, p := range m.places {
		result += p.Accept(v)
	}
	return result
}

// NewMoscow creates moscow
func NewMoscow() Moscow {
	return &moscow{
		places: []place{},
	}
}
