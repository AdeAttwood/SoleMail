package database

import (
	"testing"

	"github.com/AdeAttwood/SoleMail/pkg/assert"
)

type threadData struct {
	query  string
	match  bool
	thread *Thread
}

var threadDataProvider = []threadData{
	{
		query: "tag:\"inbox\"",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox"},
		},
	},
	{
		query: "tag:inbox",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox"},
		},
	},
	{
		query: "     tag:inbox",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox"},
		},
	},
	{
		query: "tag:inbox        ",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox"},
		},
	},
	{
		query: "     tag:inbox        ",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox"},
		},
	},
	{
		query: "tag:inbox tag:unread",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox", "unread"},
		},
	},
	{
		query: "tag:\"inbox\" tag:\"unread\"",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox", "unread"},
		},
	},
	{
		query: "tag:inbox or tag:unread",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox"},
		},
	},
	{
		query: "tag:inbox and tag:unread",
		match: false,
		thread: &Thread{
			Tags: []string{"inbox"},
		},
	},
	{
		query: "tag:inbox and (tag:unread or tag:another)",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox", "unread"},
		},
	},
	{
		query: "tag:inbox and (tag:unread or tag:another)",
		match: false,
		thread: &Thread{
			Tags: []string{"inbox", "yet another"},
		},
	},
	{
		query: "tag:inbox (tag:unread or tag:another)",
		match: false,
		thread: &Thread{
			Tags: []string{"inbox", "yet another"},
		},
	},
	{
		query: "(tag:inbox tag:unread) or (tag:one tag:two)",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox", "unread"},
		},
	},
	{
		query: "(tag:one tag:two) or (tag:inbox tag:unread)",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox", "unread"},
		},
	},
	{
		query: "(tag:one and tag:two) or (tag:inbox and tag:unread)",
		match: true,
		thread: &Thread{
			Tags: []string{"inbox", "unread"},
		},
	},
}

func TestParseQuery(t *testing.T) {
	for _, data := range threadDataProvider {
		t.Run(data.query, func(t *testing.T) {
			assert := assert.New(t)

			matched, err := MatchThread(data.query, *data.thread)
			assert.True(err == nil)

			if data.match {
				assert.True(matched)
			} else {
				assert.False(matched)
			}
		})
	}
}

func TestInvalidQuery(t *testing.T) {
	assert := assert.New(t)

	thread := Thread{
		Tags: []string{"inbox", "unread"},
	}

	_, err := MatchThread("invalid", thread)
	assert.True(err != nil)
}
