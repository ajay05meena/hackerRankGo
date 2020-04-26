package missing_numbers_test

import (
	"github.com/stretchr/testify/assert"
	. "hackerRankGo/missing_numbers"
	"testing"
)

func TestMissingNumbers(t *testing.T) {
	arr := [] int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204}
	brr := [] int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206}
	numbers := MissingNumbers(arr, brr)
	assert.Equal(t, int32(204), numbers[0])
}
