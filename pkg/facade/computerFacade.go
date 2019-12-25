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
	result := []string{
		f.cpu.Freeze(),
		f.memory.Load(f.memory.position, f.hardDrive.Read()),
		f.cpu.Jump(f.memory.position),
		f.cpu.Execute(),
	}
	return strings.Join(result, "\n")
}

// NewComputerFacade creates computerFacade
func NewComputerFacade(position string, lba, size int) *computerFacade {
	return &computerFacade{
		cpu:       NewCPU(),
		memory:    NewMemory(position),
		hardDrive: NewHardDrive(lba, size),
	}
}
