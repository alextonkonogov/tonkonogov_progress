package computerFacade

import (
	"strings"
)

// computerFacade implements facade
type computerFacade struct {
	cpu       cpu
	memory    memory
	hardDrive hardDrive
}

// Start returns how computer will start.
func (f *computerFacade) Start() string {
	result := []string{
		f.cpu.Freeze(),
		f.memory.Load(f.hardDrive.Read()),
		f.cpu.Jump(f.memory.GetPosition()),
		f.cpu.Execute(),
	}
	return strings.Join(result, "\n")
}

// cpu provides a cpu interface
type cpu interface {
	Freeze() string
	Jump(position string) string
	Execute() string
}

// memory provides a memory interface
type memory interface {
	Load(data string) string
	GetPosition() string
}

// hardDrive provides a hardDrive interface
type hardDrive interface {
	Read() string
}

// NewComputerFacade creates computerFacade
func NewComputerFacade(cpu cpu, memory memory, hardDrive hardDrive) *computerFacade {
	return &computerFacade{
		cpu:       cpu,
		memory:    memory,
		hardDrive: hardDrive,
	}
}
