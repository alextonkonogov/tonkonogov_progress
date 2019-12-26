package facade

import "fmt"

// memory implements a subsystem "memory"
type memory struct {
	sound    string
	position string
}

// Load implementation
func (m *memory) Load(position, data string) string {
	return fmt.Sprintf("Loading from %v data: '%v'.", m.position, data)
}

func (m *memory) MakeSound() string {
	return fmt.Sprintf("...%v...", m.sound)
}

// NewMemory creates memory
func NewMemory(position string) *memory {
	return &memory{
		sound:    "sounds of working memory",
		position: position,
	}
}
