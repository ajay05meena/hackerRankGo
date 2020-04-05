package mars_exploration_test

import (
	"github.com/stretchr/testify/assert"
	. "hackerRankGo/mars_exploration"
	"testing"
)

func TestMarsExploration(t *testing.T) {
	exploration := MarsExploration("SOSSOT")
	assert.Equal(t, int32(1), exploration)
}
