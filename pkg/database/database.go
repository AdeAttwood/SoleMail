package database

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

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

	it := threads_db.NewIter(nil)
	defer it.Close()
	it.Last()

	max_thread_id, err := strconv.Atoi(string(it.Key()))
	if err != nil {
		max_thread_id = 1
	}

	return &Database{
		messages:      message_db,
		threads:       threads_db,
		max_thread_id: max_thread_id,
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

func (d *Database) WriteMessage(message Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return d.messages.Set([]byte(message.MessageID), data, pebble.Sync)
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

func (d *Database) WriteThread(thread Thread) error {
	data, err := json.Marshal(thread)
	if err != nil {
		return err
	}

	return d.threads.Set([]byte(fmt.Sprint(thread.ThreadID)), data, pebble.Sync)
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

	err = d.WriteThread(t)
	return t, err
}

func (d *Database) TagQuery(query string, tag_string string) error {
	it := d.threads.NewIter(nil)
	defer it.Close()
	for it.First(); it.Valid(); it.Next() {
		thread := Thread{}
		err := json.Unmarshal(it.Value(), &thread)
		if err != nil {
			return err
		}

		matched, err := MatchThread(query, thread)
		if err != nil {
			return err
		}

		if matched {
			d.TagThread(thread.ThreadID, tag_string)
		}
	}

	return nil
}

func (d *Database) TagMessage(message_id string, tag_string string) (Message, error) {
	m, err := d.GetMessage(message_id)
	if err != nil {
		return m, err
	}

	m.Tags = ApplyTags(m.Tags, tag_string)
	err = d.WriteMessage(m)
	return m, err
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

	sort.Slice(messages, func(a, b int) bool {
		return messages[a].Date.After(messages[b].Date)
	})

	return messages, nil
}

func (d *Database) GetThreadId(message Message) int {
	for _, reference := range message.References {
		m, err := d.GetMessage(reference)
		if err == nil {
			return m.ThreadID
		}
	}

	return -1
}
