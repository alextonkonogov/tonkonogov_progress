package mother

// child provides a child interface
type child interface {
	CollectToys(direction string)
	MakeTheBed(direction string)
	VacuumTheFloor(direction string)
}

// Mother provides a Mother public interface
type Mother interface {
	Cleaning(string, string, string)
}

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

// NewMother creates mother
func NewMother(child child) Mother {
	return &mother{child: child}
}
