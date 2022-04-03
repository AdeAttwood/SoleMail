package database

import (
	"testing"

	"github.com/AdeAttwood/SoleMail/pkg/assert"
)

type tagsTestData struct {
	query  string
	input  []string
	output []string
}

var tagsTestDataProvider = []tagsTestData{
	{
		query:  "-testing",
		input:  []string{"testing"},
		output: []string{},
	},
	{
		query:  "-inbox",
		input:  []string{"inbox", "unread", "another"},
		output: []string{"unread", "another"},
	},
	{
		query:  "+another",
		input:  []string{"inbox", "unread"},
		output: []string{"inbox", "unread", "another"},
	},
	{
		query:  "-unread +another",
		input:  []string{"inbox", "unread"},
		output: []string{"inbox", "another"},
	},
}

func TestApplyTags(t *testing.T) {
	for _, data := range tagsTestDataProvider {
		t.Run(data.query, func(t *testing.T) {
			assert := assert.New(t)
			result := ApplyTags(data.input, data.query)

			assert.StringSliceLength(len(data.output), result)
		})
	}
}
