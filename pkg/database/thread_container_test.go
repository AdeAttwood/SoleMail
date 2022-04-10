package database

import (
	"testing"

	"github.com/AdeAttwood/SoleMail/pkg/assert"
)

func TestThreadContainerContains(t *testing.T) {
	assert := assert.New(t)

	a := &thread_container{}
	b := &thread_container{}
	c := &thread_container{}

	a.append(b)
	assert.True(a.contains(b))
	assert.False(a.contains(c))
}

func TestThreadContainerHeadAndTail(t *testing.T) {
	assert := assert.New(t)

	a := &thread_container{}
	b := &thread_container{}
	c := &thread_container{}

	a.append(b)
	a.append(c)

	assert.True(a.tail() == c)
	assert.True(b.head() == c.head())
	assert.True(a == c.head())
}
