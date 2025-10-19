package golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Rangga", "Dwi", "Mahendra"}
	PrintBag(names)

	assert.Equal(t, "Rangga", names[0])
	assert.Equal(t, "Dwi", names[1])
	assert.Equal(t, "Mahendra", names[2])
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{100, 200, 300, 400, 500}
	PrintBag(numbers)

	assert.Equal(t, 100, numbers[0])
	assert.Equal(t, 200, numbers[1])
	assert.Equal(t, 300, numbers[2])
	assert.Equal(t, 400, numbers[3])
	assert.Equal(t, 500, numbers[4])
}
