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
