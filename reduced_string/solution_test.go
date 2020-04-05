package reduced_string_test

import (
	"github.com/stretchr/testify/assert"
	"hackerRankGo/reduced_string"
	"testing"
)

func TestSuperReducedString(t *testing.T) {
	reducedString := reduced_string.SuperReducedString("aaabccddd")
	assert.Equal(t, "abd", reducedString, "failed")

	reducedString = reduced_string.SuperReducedString("baab")
	assert.Equal(t, "Empty String", reducedString, "failed")
}
