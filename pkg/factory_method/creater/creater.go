package creater

import (
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/hammer"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/sword"
	"github.com/alextonkonogov/tonkonogov_progress/pkg/factory_method/whip"
	"log"
)

// Weapon provides a product interface.
// All products returned by factory must provide a single interface.
type Weapon interface {
	Use() string // Every product should be can be used
}

// Blacksmith provides a factory interface.
type Blacksmith interface {
	CreateWeapon(string) Weapon // Factory Method
}

// ConcreteCreater implements Blacksmith interface.
type blacksmith struct{}

// CreateWeapon is a Factory Method
func (b *blacksmith) CreateWeapon(action string) Weapon {
	var product Weapon

	switch action {
	case "sword":
		product = sword.NewSword()
	case "hammer":
		product = hammer.NewHammer()
	case "whip":
		product = whip.NewWhip()
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

// NewBlacksmith is the ConcreteCreater constructor.
func NewBlacksmith() Blacksmith {
	return &blacksmith{}
}
