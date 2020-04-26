package coin_change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetWays(t *testing.T) {
	coins := []int64{ 3, 1, 2}
	num := 4
	ways := GetWays(int32(num), coins)
	assert.Equal(t, int64(4), ways)

	coins = []int64{8, 3, 1, 2}
	num = 3
	ways = GetWays(int32(num), coins)
	assert.Equal(t, int64(3), ways)


}
