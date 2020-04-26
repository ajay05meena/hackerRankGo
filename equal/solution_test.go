package equal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqual(t *testing.T) {
	coins := []int32{ 2, 2, 3, 7}
	equal := Equal(coins)
	assert.Equal(t, int32(2), equal)
}
