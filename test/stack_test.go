package test

import (
	adt "TDD-in-GoLang/AbstractDataTypes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackSize(t *testing.T) {
	stack := adt.NewStack()

	t.Run("initially the stack must be empty.", func(t *testing.T) {
		assert.True(t, stack.IsEmpty())
	})

	t.Run("when the stack Size Is Two", func(t *testing.T) {
		stack.Bury("Amr")
		stack.Bury("Ahmed")
		assert.Equal(t, 2, stack.Size())
	})

	t.Run("when the stack Size Is One", func(t *testing.T) {
		_, err := stack.UnBury()
		assert.Nil(t, err)
		assert.Equal(t, 1, stack.Size())
	})

	t.Run("when the stack Size Is Zero", func(t *testing.T) {
		_, err := stack.UnBury()
		assert.Nil(t, err)
		assert.Zero(t, stack.Size())
	})
}

func TestStackBury(t *testing.T) {
	stack := adt.NewStack()
	stack.Bury("Amr")
	stack.Bury("Amr")
	assert.Equal(t, 2, stack.Size())
}

func TestStackUnBury(t *testing.T) {
	stack := adt.NewStack()
	t.Run("when the stack contains no value", func(t *testing.T) {
		result, err := stack.UnBury()
		assert.NotNil(t, err.Error())
		assert.Nil(t, result)
		assert.True(t, stack.IsEmpty())
	})

	t.Run("when the stack contains at least one value", func(t *testing.T) {
		stack.Bury("Amr")
		stack.Bury("Ahmed")

		result, err := stack.UnBury()
		assert.Nil(t, err)
		assert.Equal(t, "Ahmed", result)
		assert.Equal(t, 1, stack.Size())

		result, err = stack.UnBury()
		assert.Nil(t, err)
		assert.Equal(t, "Amr", result)
		assert.True(t, stack.IsEmpty())
	})
}
