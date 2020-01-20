package main

import (
	"fmt"

	"github.com/alextonkonogov/tonkonogov_progress/pkg/command"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/command/lamp"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/command/switchOffComand"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/command/switchOnCommand"
)

func main() {
	lamp := lamp.NewLamp()
	switchOn := switchOnCommand.NewSwitchOnCommand(lamp)
	switchOff := switchOffComand.NewSwitchOffCommand(lamp)

	toggle := command.NewToggle()
	toggle.StoreCommand(switchOn)
	toggle.StoreCommand(switchOff)
	result := toggle.ExecuteCommands()
	fmt.Println(result)
}
