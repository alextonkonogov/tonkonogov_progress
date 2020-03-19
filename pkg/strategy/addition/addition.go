package addition

// Applier provides an Applier interface
type Applier interface {
	Apply(int, int) int
}

// addition implements an Applier interface
type addition struct {
}

// Apply implementation
func (a *addition) Apply(f, s int) int {
	return f + s
}

// NewAddition creates a new Applier
func NewAddition() Applier {
	return &addition{}
}
