package multiplication

// Applier provides an Applier interface
type Applier interface {
	Apply(int, int) int
}

// multiplication implements an Applier interface
type multiplication struct {
}

// Apply implementation
func (m *multiplication) Apply(f, s int) int {
	return f * s
}

// NewMultiplication creates a new Applier
func NewMultiplication() Applier {
	return &multiplication{}
}
