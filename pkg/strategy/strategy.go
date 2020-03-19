package strategy

// Applier provides a Applier interface
type Applier interface {
	Apply(int, int) int
}

// Context provides a Context interface
type Context interface {
	Action(Applier)
	Apply(int, int) int
}

// context implements a Context interface
type context struct {
	strategy Applier
}

// Action implementation
func (c *context) Action(a Applier) {
	c.strategy = a
}

// Apply implementation
func (c *context) Apply(f, s int) int {
	return c.strategy.Apply(f, s)
}

// NewContext creates a context
func NewContext() Context {
	return &context{}
}
