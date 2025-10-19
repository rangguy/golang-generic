package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (manager *MyManager) GetName() string {
	return manager.Name
}

func (manager *MyManager) GetManagerName() string {
	return manager.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Rangga", GetName[Manager](&MyManager{"Rangga"}))
	assert.Equal(t, "Dwi", GetName[VicePresident](&MyVicePresident{"Dwi"}))
}
