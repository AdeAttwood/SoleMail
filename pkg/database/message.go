package database

import "time"

type Message struct {
	MessageID  string
	ThreadID   int
	Account    string
	MessageKey string
	Subject    string
	FromName   string
	From       string
	Date       time.Time

	References []string
	Tags       []string
}

type Thread struct {
	ThreadID int
	Date     time.Time
	Subject  string
	FromName string
	From     string
	Messages []string
	Tags     []string
}

// Tests to see if this thread contains a message from a message id
func (t Thread) HasMessage(message_id string) bool {
	for _, thread_message_id := range t.Messages {
		if message_id == thread_message_id {
			return true
		}
	}

	return false
}

// Tests to see if the thread has a tag already applied.
func (t Thread) HasTag(tag string) bool {
	for _, message_tag := range t.Tags {
		if tag == message_tag {
			return true
		}
	}

	return false
}
