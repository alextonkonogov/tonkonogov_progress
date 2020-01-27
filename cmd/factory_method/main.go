package main

import (
	"fmt"

	c "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/creater"
	h "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/hammer"
	s "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/sword"
	w "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/whip"
)

var defineWeapon = func(name string) (wp c.Weapon) {
	switch name {
	case "sword":
		wp = s.NewSword()
		return
	case "hammer":
		wp = h.NewHammer()
		return
	case "whip":
		wp = w.NewWhip()
		return
	default:
	}
	return
}

func main() {
	factory := c.NewBlacksmith()
	weapons := []c.Weapon{
		factory.CreateWeapon("sword", defineWeapon),
		factory.CreateWeapon("hammer", defineWeapon),
		factory.CreateWeapon("whip", defineWeapon),
	}

	for _, weapon := range weapons {
		if weapon == nil {
			continue
		}
		fmt.Println(weapon.Use())
	}
}
