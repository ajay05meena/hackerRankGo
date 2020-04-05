package hackerrank_in_a_string_test

import (
	"github.com/stretchr/testify/assert"
	. "hackerRankGo/hackerrank_in_a_string"
	"testing"
)

func TestHackerrankInString(t *testing.T) {
	inString := HackerrankInString("hackerrank")
	assert.Equal(t, "YES", inString)

	inString = HackerrankInString("rhackerank")
	assert.Equal(t, "NO", inString)

	inString = HackerrankInString("ahankercka")
	assert.Equal(t, "NO", inString)

}


