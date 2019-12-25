package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	position, lba, size := "0x00", 100, 1024
	expect := fmt.Sprintf("Freezing processor.\nLoading from %v data: 'Some data from sector %d with size %d'.\nJumping to: %v.\nExecuting.", position, lba, size, position)
	computer := NewComputerFacade(position, lba, size)
	result := computer.Start()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
