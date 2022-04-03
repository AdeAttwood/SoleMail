// Copyright 2022 Practically.io All rights reserved
//
// Project: exam-prepper/al-g-rhythm
//
// Package to help testing in go. By default the go language dose not provide
// any assertions in the testing package. It all relies on the `t.Fatal`
// function and doing your own if statements.
//
// This package provides `Assert` functions to make tests more readable
package assert

import (
	"reflect"
	"testing"
)

// Creates a new instance of the assert package
func New(t *testing.T) *assert {
	return &assert{t: t}
}

// The internal assert type to add all the assertion function to
type assert struct{ t *testing.T }

// Asserts two string are the same
func (a *assert) StringEquals(actual string, expected string) {
	if actual != expected {
		a.t.Fatalf("Failed asserting that '%s' matches '%s'", expected, actual)
	}
}

// Asserts that a Boolean value is true
func (a *assert) True(value bool) {
	if !value {
		a.t.Fatal("Failed asserting that 'false' is true")
	}
}

// Asserts that a Boolean value is false
func (a *assert) False(value bool) {
	if value {
		a.t.Fatal("Failed asserting that 'true' is false")
	}
}

// Asserts that the length of a string slice is correct
func (a *assert) StringSliceLength(length int, value []string) {
	if len(value) != length {
		a.t.Fatalf("Failed asserting array of '%d' is '%d'", len(value), length)
	}
}

// Asserts that two interfaces are the same using "reflect.DeepEqual"
func (a *assert) DeepEqual(b, c interface{}) {
	if !reflect.DeepEqual(b, c) {
		a.t.Fatalf("Failed asserting interface '%v' is equal to '%v'", b, c)
	}
}
