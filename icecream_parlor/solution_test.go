package icecream_parlor_test

import (
	"github.com/stretchr/testify/assert"
	"hackerRankGo/icecream_parlor"
	"testing"
)

func TestIcecreamParlor(t *testing.T) {
	m := int32(6)
	arr := []int32{1, 3, 4, 5, 6}
	parlor := icecream_parlor.IcecreamParlor(m, arr)
	assert.Equal(t, int32(1), parlor[0])
	assert.Equal(t, int32(4), parlor[1])

}
