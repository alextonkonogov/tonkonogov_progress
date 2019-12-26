package facade

import (
	"fmt"
	"strings"
)

// computerFacade implements facade
type computerFacade struct {
	sound     string
	cpu       *cpu
	memory    *memory
	hardDrive *hardDrive
}

func (f *computerFacade) MakeSound() string {
	return fmt.Sprintf("...%v...", f.sound)
}

// Start returns how computer will start.
func (f *computerFacade) Start() string {
	result := []string{
		makeComponentWorkingSound(f),
		f.cpu.Freeze(),
		makeComponentWorkingSound(f.cpu),
		f.memory.Load(f.memory.position, f.hardDrive.Read()),
		makeComponentWorkingSound(f.memory),
		makeComponentWorkingSound(f.hardDrive),
		f.cpu.Jump(f.memory.position),
		f.cpu.Execute(),
	}
	return strings.Join(result, "\n")
}

// NewComputerFacade creates computerFacade
func NewComputerFacade(position string, lba, size int) *computerFacade {
	return &computerFacade{
		sound:     "sounds of starting computer",
		cpu:       NewCPU(),
		memory:    NewMemory(position),
		hardDrive: NewHardDrive(lba, size),
	}
}
