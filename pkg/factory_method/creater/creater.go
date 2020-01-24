package creater

type defineWeapon func(string) Weapon

// Weapon provides a product interface.
// All products returned by factory must provide a single interface.
type Weapon interface {
	Use() string // Every product should be can be used
}

// Blacksmith provides a factory interface.
type Blacksmith interface {
	CreateWeapon(string, defineWeapon) Weapon // Factory Method
}

// ConcreteCreater implements Blacksmith interface.
type blacksmith struct{}

// CreateWeapon is a Factory Method
func (b *blacksmith) CreateWeapon(weapon string, fn defineWeapon) Weapon {
	return fn(weapon)
}

// NewBlacksmith is the ConcreteCreater constructor.
func NewBlacksmith() Blacksmith {
	return &blacksmith{}
}
