package facade

import (
	"fmt"
	"strings"
)

//NewComputerFacade creates computerFacade
func NewComputerFacade() *computerFacade {
	return &computerFacade{
		cpu:       &cpu{},
		memory:    &memory{},
		hardDrive: &hardDrive{},
	}
}

// computerFacade implements facade
type computerFacade struct {
	cpu       *cpu
	memory    *memory
	hardDrive *hardDrive
}

//Start returns how computer will start.
func (f *computerFacade) Start() string {
	position := "0x00"
	result := []string{
		f.cpu.Freeze(),
		f.memory.Load(position, f.hardDrive.Read(100, 1024)),
		f.cpu.Jump(position),
		f.cpu.Execute(),
	}
	return strings.Join(result, "\n")
}

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

// memory implements a subsystem "memory"
type memory struct{}

// Load implementation
func (m *memory) Load(position, data string) string {
	return fmt.Sprintf("Loading from %v data: '%v'.", position, data)
}

// hardDrive implements a subsystem "hardDrive"
type hardDrive struct{}

// Read implementation
func (h *hardDrive) Read(lba, size int) string {
	return fmt.Sprintf("Some data from sector %d with size %d", lba, size)
}
