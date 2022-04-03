package database

import "errors"

type boolstack []bool

// IsEmpty: check if stack is empty
func (s *boolstack) isEmpty() bool {
	return len(*s) == 0
}

func (s *boolstack) length() int {
	return len(*s)
}

// Push a new value onto the stack
func (s *boolstack) push(item bool) {
	*s = append(*s, item)
}

// Remove and return top element of stack. Return an error if the stack is empty
func (s *boolstack) pop() (bool, error) {
	if s.isEmpty() {
		return false, errors.New("Stack is empty")
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, nil
	}
}
