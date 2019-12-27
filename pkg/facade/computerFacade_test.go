package facade

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	position string
	lba      int
	size     int
}{
	{"0x00", 62, 128},
	{"0x2a", 63, 256},
	{"0b10", 945, 512},
	{"0x20", 1007, 1024},
	{"0x3e", 1070, 2048},
	{"0x5b", 2015, 4096},
}

func TestFacade(t *testing.T) {
	for i, tt := range tests {
		t.Run(tt.position, func(t *testing.T) {
			position, lba, size := tt.position, tt.lba, tt.size
			expect := fmt.Sprintf("...sounds of starting computer...\nFreezing processor.\n...sounds of working cpu...\nLoading from %v data: 'Some data from sector %d with size %d'.\n...sounds of working memory...\n...sounds of working hard drive...\nJumping to: %v.\nExecuting.", position, lba, size, position)
			computer := NewComputerFacade(position, lba, size)
			result := computer.Start()

			assert.EqualValues(t, expect, result, "Result is not equal to the expected one")
			t.Log("test №", i)
		})
	}
}
