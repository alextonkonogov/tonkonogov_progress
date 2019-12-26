package facade

type sounder interface {
	MakeSound() string
}

func makeComponentWorkingSound(s sounder) string {
	return s.MakeSound()
}
