package main

import (
	"fmt"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/facade/computerFacade"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/facade/cpu"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/facade/hardDrive"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/facade/memory"
)

var position, lba, size = "0x01", 200, 2048

func main() {
	computer :=
		computerFacade.NewComputerFacade(
			cpu.NewCPU(),
			memory.NewMemory(position),
			hardDrive.NewHardDrive(lba, size))
	result := computer.Start()
	fmt.Println(result)
}
