package memory

import "fmt"

// Memory provides a Memory interface
type Memory interface {
	Load(data string) string
	GetPosition() string
}

// memory implements a subsystem "memory"
type memory struct {
	position string
}

// Load implementation
func (m *memory) Load(data string) string {
	return fmt.Sprintf("Loading from %v data: '%v'.", m.position, data)
}

// GetPosition implementation
func (m *memory) GetPosition() string {
	return m.position
}

// NewMemory creates memory
func NewMemory(position string) Memory {
	return &memory{
		position: position,
	}
}
