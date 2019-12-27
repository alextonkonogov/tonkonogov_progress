package builder

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	toys, bed, floor string
}{
	{"all", "accurately", "carefully"},
	{"scattered", "properly", "normally"},
	{"damned", "quickly", "completely"},
}

func TestBuilder(t *testing.T) {
	for i, tt := range tests {
		name := strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			toys, bed, floor := tt.toys, tt.bed, tt.floor
			expect := fmt.Sprintf("Collected %v toys\nMade the bed %v\nVacuumed the floor %v\n", toys, bed, floor)
			order := NewOrder()
			mother := NewMother(NewRoomCleaning(order))
			mother.Cleaning(toys, bed, floor)
			result := order.Show()

			fmt.Println(result)
			assert.EqualValues(t, expect, result, "Result is not equal to the expected one")
		})
	}
}
