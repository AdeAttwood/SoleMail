package assert

import (
	"testing"
)

func TestAssertStringEquals(t *testing.T) {
	assert := New(t)
	assert.StringEquals("One", "One")
}

func TestAssertTrue(t *testing.T) {
	assert := New(t)
	assert.True(true)
}

func TestAssertFalse(t *testing.T) {
	assert := New(t)
	assert.False(false)
}

func TestAssertStringSlice(t *testing.T) {
	assert := New(t)
	assert.StringSliceLength(2, []string{"one", "two"})
}
