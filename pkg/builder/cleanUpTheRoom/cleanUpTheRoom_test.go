package cleanUpTheRoom

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	motherT "github.com/alextonkonogov/tonkonogov_progress/pkg/builder/mother"
	orderT "github.com/alextonkonogov/tonkonogov_progress/pkg/builder/order"
)

type data struct {
	toys, bed, floor string
}

var tests = []data{
	{"all", "accurately", "carefully"},
	{"scattered", "properly", "normally"},
	{"damned", "quickly", "completely"},
}

func expect(d data) string {
	return fmt.Sprintf("Collected %v toys\nMade the bed %v\nVacuumed the floor %v\n", d.toys, d.bed, d.floor)
}

func TestBuilder(t *testing.T) {
	for i, tt := range tests {
		name := strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			toys, bed, floor := tt.toys, tt.bed, tt.floor
			expect := expect(tt)
			order := orderT.NewOrder()
			mother := motherT.NewMother(NewRoomCleaning(order))
			mother.Cleaning(toys, bed, floor)
			result := order.Show()
			assert.EqualValues(t, expect, result, "Result is not equal to the expected one")
		})
	}
}
