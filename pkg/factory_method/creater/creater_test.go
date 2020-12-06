package creater

import (
	"testing"

	h "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/hammer"
	s "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/sword"
	w "github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/whip"
)

var assert = []string{"shuh shuh shuh!", "bang!", "whoosh whoosh!"}

var fn = func(name string) (wp Weapon) {
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

func TestFactoryMethod(t *testing.T) {
	factory := NewBlacksmith()
	products := []Weapon{
		factory.CreateWeapon("sword", fn),
		factory.CreateWeapon("hammer", fn),
		factory.CreateWeapon("whip", fn),
	}

	for i, product := range products {
		if action := product.Use(); action != assert[i] {
			t.Errorf("Expect action to %s, but %s.\n", assert[i], action)
		}
	}
}
