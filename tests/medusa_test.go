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

func TestPersonIsObject(t *testing.T) {
	person := exercises.Person{}

	assert.Equal(t, person, exercises.Person{})
}

func TestNewPerson(t *testing.T) {
	person := exercises.NewPerson("Jhun")

	assert.Equal(t, "Jhun", person.Name)
	assert.False(t, person.Stoned, "Person is not stoned")
}

func TestMedusaStare(t *testing.T) {
	medusa := exercises.NewMedusa("Brenna")
	person := exercises.NewPerson("Jhun")

	assert.Equal(t, []string{}, medusa.Statues)
	assert.False(t, person.Stoned, "Person is not stoned")

	medusa.Stare(&person)

	assert.True(t, person.Stoned, "Jhun is stoned")
	assert.Equal(t, 1, len(medusa.Statues))
}
