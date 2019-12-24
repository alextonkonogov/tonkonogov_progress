package main

import (
	"fmt"
	"tonkonogov_progress/pkg/facade"
)

func main() {
	computer := facade.NewComputerFacade()
	result := computer.Start()
	fmt.Println(result)
}
