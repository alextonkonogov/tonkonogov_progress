package creater

import (
	"testing"
)

var assert = []string{"shuh shuh!", "bang!", "whoosh whoosh!"}

func TestFactoryMethod(t *testing.T) {
	factory := NewBlacksmith()
	products := []Weapon{
		factory.CreateWeapon("sword"),
		factory.CreateWeapon("hammer"),
		factory.CreateWeapon("whip"),
	}

	for i, product := range products {
		if action := product.Use(); action != assert[i] {
			t.Errorf("Expect action to %s, but %s.\n", assert[i], action)
		}
	}
}
