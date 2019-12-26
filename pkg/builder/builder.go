package builder

// child provides a child interface
type child interface {
	CollectToys(direction string)
	MakeTheBed(direction string)
	VacuumTheFloor(direction string)
}
