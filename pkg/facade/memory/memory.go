package memory

import "fmt"

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
func NewMemory(position string) *memory {
	return &memory{
		position: position,
	}
}
