package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	toys, bed, floor := "all", "accurately", "carefully"
	expect := fmt.Sprintf("Collected %v toys\nMade the bed %v\nVacuumed the floor %v\n", toys, bed, floor)

	order := NewOrder()
	mother := NewMother(NewCleanUpTheRoom(order))
	mother.Cleaning(toys, bed, floor)
	result := order.Show()

	if result != expect {
		t.Errorf("Expect result to %s, but %s", result, expect)
	}
}
