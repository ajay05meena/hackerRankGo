package closest_numbers_test

import (
	"github.com/stretchr/testify/assert"
	. "hackerRankGo/closest_numbers"
	"testing"
)

func TestClosestNumbers(t *testing.T) {
	a := make([]int32, 1, 1)
	a = append(a, 5, 4, 3, 2)
	numbers := ClosestNumbers(a)
	assert.Equal(t, 6, len(numbers))
}
