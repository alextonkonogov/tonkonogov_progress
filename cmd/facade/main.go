package main

import (
	"fmt"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/facade"
)

var position, lba, size = "0x01", 200, 2048

func main() {
	computer := facade.NewComputerFacade(position, lba, size)
	result := computer.Start()
	fmt.Println(result)
}
