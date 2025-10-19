package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	}

	return false
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[int](1, 1))
	assert.True(t, IsSame[string]("rangga", "rangga"))
	assert.False(t, IsSame[string]("rangga", "Rangga"))
}
