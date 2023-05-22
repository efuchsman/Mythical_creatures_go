package tests

import (
	"testing"

	"example.com/mythical_creatures_go/exercises"
	"github.com/stretchr/testify/assert"
)

func TestMedusaIsFunction(t *testing.T) {
	medusa := exercises.Medusa{}
	assert.IsType(t, exercises.Medusa{}, medusa, "medusa should be a function")
}

func TestMedusaIsObject(t *testing.T) {
	medusa := exercises.Medusa{}

	assert.Equal(t, medusa, exercises.Medusa{})
}

func TestNewMedusa(t *testing.T) {
	medusa := exercises.NewMedusa("Bree")

	assert.Equal(t, "Bree", medusa.Name)
	assert.Equal(t, []string{}, medusa.Statues)
}
