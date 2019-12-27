package builder

import "fmt"

// mother implements a mother (manager)
type mother struct {
	son child
}

// Cleaning tells the child (doer) what to do, in what order and how
func (m *mother) Cleaning(toys, bed, floor string) {
	m.son.CollectToys(toys)
	m.son.MakeTheBed(bed)
	m.son.VacuumTheFloor(floor)
}

func (m *mother) Call() {
	fmt.Println("...mother is initiated")
}

// NewMother creates mother
func NewMother(son child) *mother {
	return &mother{son: son}
}
