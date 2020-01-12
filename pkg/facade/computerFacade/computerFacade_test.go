package computerFacade

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	cpuT "github.com/alextonkonogov/tonkonogov_progress/pkg/facade/cpu"
	hardDriveT "github.com/alextonkonogov/tonkonogov_progress/pkg/facade/hardDrive"
	memoryT "github.com/alextonkonogov/tonkonogov_progress/pkg/facade/memory"
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
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			position, lba, size := tt.position, tt.lba, tt.size
			expect := fmt.Sprintf("Freezing processor.\nLoading from %v data: 'Some data from sector %d with size %d'.\nJumping to: %v.\nExecuting.", position, lba, size, position)
			computer :=
				NewComputerFacade(
					cpuT.NewCPU(),
					memoryT.NewMemory(position),
					hardDriveT.NewHardDrive(lba, size))
			result := computer.Start()

			assert.EqualValues(t, expect, result, "Result is not equal to the expected one")
		})
	}
}
