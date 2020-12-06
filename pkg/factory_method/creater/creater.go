package creater

// defineWeapon implements a function returning a Weapon
type defineWeapon func(string) Weapon

// Weapon provides a Weapon interface
type Weapon interface {
	Use() string
}

// Blacksmith provides a Blacksmith (factory) interface
type Blacksmith interface {
	CreateWeapon(string, defineWeapon) Weapon // Factory Method
}

// blacksmith implements a blacksmith interface
type blacksmith struct{}

// CreateWeapon implementation
func (b *blacksmith) CreateWeapon(weapon string, fn defineWeapon) Weapon {
	return fn(weapon)
}

// NewBlacksmith creates a Blacksmith
func NewBlacksmith() Blacksmith {
	return &blacksmith{}
}
