package cpu

import "fmt"

// cpu implements a subsystem "cpu"
type cpu struct {
}

// Freeze implementation
func (c *cpu) Freeze() string {
	return fmt.Sprint("Freezing processor.")
}

// Jump implementation
func (c *cpu) Jump(position string) string {
	return fmt.Sprintf("Jumping to: %v.", position)
}

// Execute implementation
func (c *cpu) Execute() string {
	return fmt.Sprint("Executing.")
}

// NewCPU creates cpu
func NewCPU() *cpu {
	return &cpu{}
}
