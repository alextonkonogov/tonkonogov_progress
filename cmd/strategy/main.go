package main

import (
	"fmt"

	c "github.com/alextonkonogov/tonkonogov_progress/pkg/strategy"
	a "github.com/alextonkonogov/tonkonogov_progress/pkg/strategy/addition"
	m "github.com/alextonkonogov/tonkonogov_progress/pkg/strategy/multiplication"
)

var first, second = 4, 6

func main() {
	ctx := c.NewContext()

	ctx.Action(m.NewMultiplication())
	fmt.Println(ctx.Apply(first, second))

	ctx.Action(a.NewAddition())
	fmt.Println(ctx.Apply(first, second))
}
