package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"testing"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Rangga",
	}

	second := map[string]string{
		"Name": "Rangga",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlices(t *testing.T) {
	first := []int{1, 2, 3}
	second := []int{1, 2, 3}

	assert.True(t, slices.Equal(first, second))
}
