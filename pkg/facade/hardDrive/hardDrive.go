package hardDrive

import "fmt"

// HardDrive provides a HardDrive interface
type HardDrive interface {
	Read() string
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
func NewHardDrive(lba, size int) HardDrive {
	return &hardDrive{
		lba:  lba,
		size: size,
	}
}
