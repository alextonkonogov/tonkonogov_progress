package facade

import "fmt"

// cpu implements a subsystem "cpu"
type cpu struct{}

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

// memory implements a subsystem "memory"
type memory struct {
	position string
}

// Load implementation
func (m *memory) Load(position, data string) string {
	return fmt.Sprintf("Loading from %v data: '%v'.", m.position, data)
}

// NewMemory creates memory
func NewMemory(position string) *memory {
	return &memory{
		position: position,
	}
}

// hardDrive implements a subsystem "hardDrive"
type hardDrive struct {
	lba, size int
}

// Read implementation
func (h *hardDrive) Read() string {
	return fmt.Sprintf("Some data from sector %d with size %d", h.lba, h.size)
}

// NewHardDrive creates hardDrive
func NewHardDrive(lba, size int) *hardDrive {
	return &hardDrive{
		lba:  lba,
		size: size,
	}
}
