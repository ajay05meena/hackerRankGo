package picking_numbers_test

import (
	. "hackerRankGo/picking_numbers"
	"testing"
)

func Test_PickingNumbers(t *testing.T) {
	a := make([]int32, 2, 5)
	a = append(a, 4, 6, 5, 3, 3, 1)
	result := PickingNumbers(a)
	if result != 3 {
		t.Error("Failed result", result)
	}
}
