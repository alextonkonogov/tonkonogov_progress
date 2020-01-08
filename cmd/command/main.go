package main

import (
	"fmt"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/command"
)

func main() {
	lamp := command.NewLamp()

	switchOn := command.NewSwitchOnCommand(lamp)
	switchOff := command.NewSwitchOffCommand(lamp)

	toggle := command.NewToggle()
	toggle.StoreCommand(switchOn)
	toggle.StoreCommand(switchOff)
	result := toggle.ExecuteCommands()
	fmt.Println(result)
}
