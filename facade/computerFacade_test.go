package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	expect := "Freezing processor.\nLoading from 0x00 data: 'Some data from sector 100 with size 1024'.\nJumping to: 0x00.\nExecuting."
	computer := NewComputerFacade()
	result := computer.Start()

	fmt.Println(result)

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
