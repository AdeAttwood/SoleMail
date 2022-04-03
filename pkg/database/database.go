package database

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/cockroachdb/pebble"
)

type Database struct {
	messages      *pebble.DB
	threads       *pebble.DB
	max_thread_id int
}

func Open(path string) (*Database, error) {
	messages_path := fmt.Sprintf("%s/db/messages", path)
	message_db, err := pebble.Open(messages_path, &pebble.Options{})
	if err != nil {
		return &Database{}, err
	}

	threads_path := fmt.Sprintf("%s/db/threads", path)
	threads_db, err := pebble.Open(threads_path, &pebble.Options{})
	if err != nil {
		return &Database{}, err
	}

	return &Database{
		messages:      message_db,
		threads:       threads_db,
		max_thread_id: 1,
	}, nil
}

func (d *Database) NextThreadID() int {
	d.max_thread_id = d.max_thread_id + 1
	return d.max_thread_id
}

func (d *Database) Close() error {
	var err error
	if err = d.messages.Close(); err != nil {
		return err
	}

	return err
}

func (d *Database) List() ([]Message, error) {
	messages := []Message{}
	it := d.messages.NewIter(nil)
	defer it.Close()
	for it.First(); it.Valid(); it.Next() {
		message := Message{}
		err := json.Unmarshal(it.Value(), &message)
		if err != nil {
			return messages, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (d *Database) ListThreads(query string) ([]Thread, error) {
	threads := []Thread{}
	it := d.threads.NewIter(nil)
	defer it.Close()
	for it.First(); it.Valid(); it.Next() {
		thread := Thread{}
		err := json.Unmarshal(it.Value(), &thread)
		if err != nil {
			return threads, err
		}

		if len(query) == 0 {
			threads = append(threads, thread)
			continue
		}

		matched, err := MatchThread(query, thread)
		if err != nil {
			return threads, err
		}

		if matched {
			threads = append(threads, thread)
		}
	}

	sort.Slice(threads, func(a, b int) bool {
		return threads[a].Date.After(threads[b].Date)
	})

	return threads, nil
}

func (d *Database) GetMessage(message_id string) (Message, error) {
	message := Message{}
	data, closer, err := d.messages.Get([]byte(message_id))
	if err != nil {
		return message, err
	}

	defer closer.Close()

	err = json.Unmarshal(data, &message)
	if err != nil {
		return message, err
	}

	return message, nil
}

func (d *Database) GetThread(thread_id int) (Thread, error) {
	thread := Thread{}
	data, closer, err := d.threads.Get([]byte(fmt.Sprint(thread_id)))
	if err != nil {
		return thread, err
	}

	defer closer.Close()

	err = json.Unmarshal(data, &thread)
	if err != nil {
		return thread, err
	}

	return thread, nil
}

func (d *Database) TagThread(thread_id int, tag_string string) (Thread, error) {
	t, err := d.GetThread(thread_id)
	if err != nil {
		return t, err
	}

	t.Tags = ApplyTags(t.Tags, tag_string)
	for _, message_id := range t.Messages {
		_, err := d.TagMessage(message_id, tag_string)
		if err != nil {
			return t, err
		}
	}

	data, err := json.Marshal(t)
	if err != nil {
		return t, err
	}

	d.threads.Set([]byte(fmt.Sprint(t.ThreadID)), data, pebble.Sync)

	return t, nil
}

func (d *Database) TagMessage(message_id string, tag_string string) (Message, error) {
	m, err := d.GetMessage(message_id)
	if err != nil {
		return m, err
	}

	m.Tags = ApplyTags(m.Tags, tag_string)
	data, err := json.Marshal(m)
	if err != nil {
		return m, err
	}

	d.messages.Set([]byte(m.MessageID), data, pebble.Sync)

	return m, nil
}

func (d *Database) ReadThread(thread_id int) ([]Message, error) {
	messages := []Message{}
	thread, err := d.GetThread(thread_id)
	if err != nil {
		return messages, err
	}

	for _, message_id := range thread.Messages {
		message, err := d.GetMessage(message_id)
		if err != nil {
			return messages, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (d *Database) AddMessage(message Message) error {
	if len(message.References) == 0 {
		message.ThreadID = d.NextThreadID()
	} else {
		m, err := d.GetMessage(message.References[0])
		if err != nil {
			return err
		}

		message.ThreadID = m.ThreadID
		message.Tags = m.Tags
	}

	m, err := d.GetMessage(message.MessageID)
	if err == nil {
		message.Tags = m.Tags
	} else {
		message.Tags = append(message.Tags, "inbox", "unread")
	}

	t, err := d.GetThread(message.ThreadID)
	if err != nil {
		t = Thread{
			ThreadID: message.ThreadID,
			Messages: []string{message.MessageID},
		}

	}

	latest_message := message
	for _, thread_message_id := range t.Messages {
		if thread_message_id == message.MessageID {
			continue
		}

		thread_message, err := d.GetMessage(thread_message_id)
		if err != nil {
			return err
		}

		if thread_message.Date.After(latest_message.Date) {
			latest_message = thread_message
		}
	}

	t.Date = latest_message.Date
	t.Subject = latest_message.Subject
	t.FromName = latest_message.FromName
	t.From = latest_message.From
	t.Tags = latest_message.Tags

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	d.threads.Set([]byte(fmt.Sprint(message.ThreadID)), data, pebble.Sync)

	data, err = json.Marshal(message)
	if err != nil {
		return err
	}

	return d.messages.Set([]byte(message.MessageID), data, pebble.Sync)
}

func (d *Database) UpdateThreads() error {
	it := d.messages.NewIter(nil)
	defer it.Close()
	for it.First(); it.Valid(); it.Next() {
		message := Message{}
		if err := json.Unmarshal(it.Value(), &message); err != nil {
			return err
		}

		thread, err := d.GetThread(message.ThreadID)
		if err != nil {
			return err
		}

		for _, message_id := range message.References {
			if !thread.HasMessage(message_id) {
				thread.Messages = append(thread.Messages, message_id)
			}
		}

		data, err := json.Marshal(thread)
		if err != nil {
			return err
		}

		d.threads.Set([]byte(fmt.Sprint(message.ThreadID)), data, pebble.Sync)
	}

	return nil
}
