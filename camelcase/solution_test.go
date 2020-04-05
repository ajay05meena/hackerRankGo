package camelcase_test

import (
	"github.com/stretchr/testify/assert"
	"hackerRankGo/camelcase"
	"testing"
)

func TestCamelcase(t *testing.T) {
	i := camelcase.Camelcase("helloWorld")
	assert.Equal(t, int32(2), i, "failed")

	i = camelcase.Camelcase("hello")
	assert.Equal(t, int32(1), i, "failed")

}
