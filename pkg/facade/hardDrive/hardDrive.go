package hardDrive

import "fmt"

// hardDrive implements a subsystem "hardDrive"
type hardDrive struct {
	lba, size int
}

// Read implementation
func (h *hardDrive) Read() string {
	return fmt.Sprintf("Some data from sector %d with size %d", h.lba, h.size)
}

// NewHardDrive creates hardDrive
func NewHardDrive(lba, size int) *hardDrive {
	return &hardDrive{
		lba:  lba,
		size: size,
	}
}
