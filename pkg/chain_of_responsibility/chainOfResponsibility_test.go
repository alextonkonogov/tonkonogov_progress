package chain_of_responsibility

import (
	a "github.com/alextonkonogov/tonkonogov_progress/pkg/chain_of_responsibility/concreteHandlerA"
	b "github.com/alextonkonogov/tonkonogov_progress/pkg/chain_of_responsibility/concreteHandlerB"
	c "github.com/alextonkonogov/tonkonogov_progress/pkg/chain_of_responsibility/concreteHandlerC"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var m = map[int]string{
	-1: "...silence...",
	0:  "...silence...",
	1:  "Im handler A",
	2:  "Im handler B",
	3:  "Im handler C",
	4:  "...silence...",
	5:  "...silence...",
	99: "...silence...",
}

func TestChainOfResponsibility(t *testing.T) {
	for k, message := range m {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			expect := message
			handlers := a.NewConcreteHandlerA(b.NewConcreteHandlerB(c.NewConcreteHandlerC()))
			result := handlers.SendRequest(k)
			assert.EqualValues(t, expect, result, "Result is not equal to the expected one")
		})
	}
}
