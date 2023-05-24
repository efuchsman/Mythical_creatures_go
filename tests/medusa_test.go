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
	assert.Equal(t, []exercises.Person{}, medusa.Statues)
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

	assert.Equal(t, []exercises.Person{}, medusa.Statues)
	assert.False(t, person.Stoned, "Person is not stoned")

	medusa.Stare(&person)

	assert.True(t, person.Stoned, "Jhun is stoned")
	assert.Equal(t, 1, len(medusa.Statues))
}

func TestMedusaStareCount(t *testing.T) {
	medusa := exercises.NewMedusa("Meeka")
	person1 := exercises.NewPerson("Louisa")
	person2 := exercises.NewPerson("Bree")
	person3 := exercises.NewPerson("TayTay")
	person4 := exercises.NewPerson("Steve")

	medusa.Stare(&person1)
	medusa.Stare(&person2)
	medusa.Stare(&person3)

	assert.Equal(t, 3, len(medusa.Statues))
	assert.True(t, person1.Stoned, "Louisa is stoned")
	assert.True(t, person2.Stoned, "Bree is stoned")
	assert.True(t, person3.Stoned, "TayTay is stoned")
	assert.False(t, person4.Stoned, "Steve is not stoned")
	// assert.Equal(t, []exercises.Person{person1, person2, person3}, medusa.Statues)

	// medusa.Stare(&person4)

	// assert.Equal(t, 3, len(medusa.Statues))
	// assert.True(t, person4.Stoned, "Steve is stoned")
	// assert.Equal(t, []exercises.Person{person1, person2, person3}, medusa.Statues)
	// assert.False(t, person1.Stoned, "Louisa is not stoned")
}
