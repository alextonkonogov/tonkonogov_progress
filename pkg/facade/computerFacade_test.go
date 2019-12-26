package facade

import (
	"fmt"
	"testing"
)

var position, lba, size = "0x00", 100, 1024

func TestFacade(t *testing.T) {
	expect :=
		fmt.Sprintf("Freezing processor.\n...sounds of working cpu...\nLoading from %v data: 'Some data from sector %d with size %d'.\n...sounds of working memory...\n...sounds of working hard drive...\nJumping to: %v.\nExecuting.",
			position, lba, size, position)
	computer := NewComputerFacade(position, lba, size)
	result := computer.Start()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
