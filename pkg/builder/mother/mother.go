package mother

// mother implements a mother (manager)
type mother struct {
	child child
}

// Cleaning tells the child (doer) what to do, in what order and how
func (m *mother) Cleaning(toys, bed, floor string) {
	m.child.CollectToys(toys)
	m.child.MakeTheBed(bed)
	m.child.VacuumTheFloor(floor)
}

// child provides a child interface
type child interface {
	CollectToys(direction string)
	MakeTheBed(direction string)
	VacuumTheFloor(direction string)
}

// NewMother creates mother
func NewMother(child child) *mother {
	return &mother{child: child}
}
