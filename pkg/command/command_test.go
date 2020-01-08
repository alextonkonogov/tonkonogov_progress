package command

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	expect := "The light is on\nThe light is off\n"

	lamp := NewLamp()

	switchOn := NewSwitchOnCommand(lamp)
	switchOff := NewSwitchOffCommand(lamp)

	toggle := NewToggle()
	toggle.StoreCommand(switchOn)
	toggle.StoreCommand(switchOff)
	result := toggle.ExecuteCommands()
	fmt.Println(result)

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
