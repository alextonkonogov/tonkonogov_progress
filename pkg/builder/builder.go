package builder

// child provides a child interface
type child interface {
	CollectToys(direction string)
	MakeTheBed(direction string)
	VacuumTheFloor(direction string)
}

// caller provides a caller interface
type caller interface {
	Call()
}

// function receives a caller interface
func Call(c caller) {
	c.Call()
}
