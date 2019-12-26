package facade

import "fmt"

// cpu implements a subsystem "cpu"
type cpu struct {
	sound string
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

func (c *cpu) MakeSound() string {
	return fmt.Sprintf("...%v...", c.sound)
}

// NewCPU creates cpu
func NewCPU() *cpu {
	return &cpu{
		sound: "sounds of working cpu",
	}
}
