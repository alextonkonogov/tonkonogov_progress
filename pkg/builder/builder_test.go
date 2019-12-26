package builder

import (
	"fmt"
	"testing"
)

var toys, bed, floor = "all", "accurately", "carefully"

func TestBuilder(t *testing.T) {
	expect := fmt.Sprintf("Collected %v toys\nMade the bed %v\nVacuumed the floor %v\n", toys, bed, floor)
	order := NewOrder()
	mother := NewMother(NewCleanUpTheRoom(order))
	mother.Cleaning(toys, bed, floor)
	result := order.Show()

	if result != expect {
		t.Errorf("Expect result to %s, but %s", result, expect)
	}
}
