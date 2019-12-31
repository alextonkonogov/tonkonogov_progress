package facade

// sounder provides a sounder interface
type sounder interface {
	MakeSound() string
}

// function receives a sounder interface
func makeComponentWorkingSound(s sounder) string {
	return s.MakeSound()
}
