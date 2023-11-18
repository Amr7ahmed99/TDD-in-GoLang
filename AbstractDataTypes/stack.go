package adt

import (
	"errors"
)

type Stack struct {
	values []string
}

func NewStack() *Stack {
	return &Stack{values: make([]string, 0)}
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.values) == 0
}

func (stack *Stack) Size() int {
	return len(stack.values)
}

func (stack *Stack) Bury(value string) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) UnBury() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("error: The unbury operation failed, the stack contains no value")
	}
	lastValue := stack.values[stack.Size()-1]
	stack.values = stack.values[0 : stack.Size()-1]
	return lastValue, nil
}
