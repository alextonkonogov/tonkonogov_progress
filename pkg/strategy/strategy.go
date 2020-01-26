package strategy

// StrategyAction provides a StrategyAction interface
type StrategyAction interface {
	Apply(int, int) int
}

// Context provides a Context interface
type Context interface {
	Action(StrategyAction)
	Apply(int, int) int
}

// context provides a context for execution of a strategy.
type context struct {
	strategy StrategyAction
}

// Action implementation
func (c *context) Action(a StrategyAction) {
	c.strategy = a
}

// Apply implementation
func (c *context) Apply(f, s int) int {
	return c.strategy.Apply(f, s)
}

// NewContext creates a new Context
func NewContext() Context {
	return &context{}
}
