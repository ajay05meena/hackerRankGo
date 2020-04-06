package find_the_median_test

import (
	"github.com/stretchr/testify/assert"
	. "hackerRankGo/find_the_median"
	"testing"
)

func TestFindMedian(t *testing.T) {
	arr := []int32{0, 1, 2, 4, 6, 5, 3}
	median := FindMedian(arr)
	assert.Equal(t, int32(3), median)
}
