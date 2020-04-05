package save_the_prisoner_test

import (
	"github.com/stretchr/testify/assert"
	. "hackerRankGo/save_the_prisoner"
	"testing"
)

func TestSaveThePrisoner(t *testing.T) {
	prisoner := SaveThePrisoner(7, 19, 2)
	assert.Equal(t, int32(6), prisoner, "failed")

	prisoner = SaveThePrisoner(7, 19, 3)
	assert.Equal(t, int32(7), prisoner, "failed")
}
