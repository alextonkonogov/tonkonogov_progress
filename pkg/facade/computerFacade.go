package facade

import (
	"strings"
)

// computerFacade implements facade
type computerFacade struct {
	cpu       *cpu
	memory    *memory
	hardDrive *hardDrive
}

// Start returns how computer will start.
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

// NewComputerFacade creates computerFacade
func NewComputerFacade() *computerFacade {
	return &computerFacade{
		cpu:       NewCPU(),
		memory:    NewMemory(),
		hardDrive: NewHardDrive(),
	}
}
