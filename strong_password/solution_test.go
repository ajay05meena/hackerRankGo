package strong_password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumNumber(t *testing.T) {
	number := MinimumNumber(3, "Ab1")
	assert.Equal(t, int32(3), number)
}
