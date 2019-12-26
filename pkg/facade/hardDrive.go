package facade

import "fmt"

// hardDrive implements a subsystem "hardDrive"
type hardDrive struct {
	sound     string
	lba, size int
}

// Read implementation
func (h *hardDrive) Read() string {
	return fmt.Sprintf("Some data from sector %d with size %d", h.lba, h.size)
}

func (h *hardDrive) MakeSound() string {
	return fmt.Sprintf("...%v...", h.sound)
}

// NewHardDrive creates hardDrive
func NewHardDrive(lba, size int) *hardDrive {
	return &hardDrive{
		sound: "sounds of working hard drive",
		lba:   lba,
		size:  size,
	}
}
