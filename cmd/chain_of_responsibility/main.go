package main

import (
	"fmt"
	a "github.com/alextonkonogov/tonkonogov_progress/pkg/chain_of_responsibility/concreteHandlerA"
	b "github.com/alextonkonogov/tonkonogov_progress/pkg/chain_of_responsibility/concreteHandlerB"
	c "github.com/alextonkonogov/tonkonogov_progress/pkg/chain_of_responsibility/concreteHandlerC"
)

func main() {
	handlers := a.NewConcreteHandlerA(b.NewConcreteHandlerB(c.NewConcreteHandlerC()))
	result := handlers.SendRequest(3)
	fmt.Println(result)
}
