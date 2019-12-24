package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {

	//expect := "Build house\nTree grow\nChild born"

	computer := NewComputerFacade()

	computer.Start()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
