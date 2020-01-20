package command

import (
	"fmt"
	"testing"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/command/lamp"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/command/switchOffComand"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/command/switchOnCommand"
)

var expect = "The light is on\nThe light is off\n"

func TestCommand(t *testing.T) {
	lamp := lamp.NewLamp()

	switchOn := switchOnCommand.NewSwitchOnCommand(lamp)
	switchOff := switchOffComand.NewSwitchOffCommand(lamp)

	toggle := NewToggle()
	toggle.StoreCommand(switchOn)
	toggle.StoreCommand(switchOff)
	result := toggle.ExecuteCommands()
	fmt.Println(result)

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
