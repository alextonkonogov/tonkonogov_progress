package main

import (
	"fmt"
	"tonkonogov_progress/pkg/facade"
)

func main() {
	position, lba, size := "0x01", 200, 2048
	computer := facade.NewComputerFacade(position, lba, size)
	result := computer.Start()
	fmt.Println(result)
}
