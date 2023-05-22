package tests

import (
	"testing"

	"example.com/mythical_creatures_go/exercises"
	"github.com/stretchr/testify/assert"
)

func TestUnicornIsFunction(t *testing.T) {
	unicorn := exercises.Unicorn{}
	assert.IsType(t, exercises.Unicorn{}, unicorn, "Unicorn should be a function")
}

func TestUnicornIsObject(t *testing.T) {
	unicorn := exercises.Unicorn{}

	assert.Equal(t, unicorn, exercises.Unicorn{})
}

func TestNewUnicorn(t *testing.T) {
	unicorn1 := exercises.NewUnicorn("Bree", "")
	unicorn2 := exercises.NewUnicorn("Pam", "")

	assert.Equal(t, "Bree", unicorn1.Name)
	assert.Equal(t, "Pam", unicorn2.Name)
}

func TestNewUnicornColor(t *testing.T) {
	unicorn1 := exercises.NewUnicorn("TayTay", "blue")

	assert.Equal(t, "blue", unicorn1.Color)
}

func TestNewUnicornDefaultColorIsWhite(t *testing.T) {
	unicorn1 := exercises.NewUnicorn("Louisa", "")

	assert.Equal(t, "white", unicorn1.Color)
}

func TestUnicornIsWhite(t *testing.T) {
	unicorn1 := exercises.NewUnicorn("Louisa", "")
	unicorn2 := exercises.NewUnicorn("Steve", "red")

	assert.True(t, unicorn1.IsWhite(), "Unicorn is white")
	assert.False(t, unicorn2.IsWhite(), "Unicorn is not white")
}

func TestUnicornSays(t *testing.T) {
	unicorn := exercises.NewUnicorn("Brenna", "")

	assert.Equal(t, "**;* Wonderful! *;**", unicorn.Says("Wonderful!"))
	assert.Equal(t, "**;* Batman has a great smile *;**", unicorn.Says("Batman has a great smile"))
}
