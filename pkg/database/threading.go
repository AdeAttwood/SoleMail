// Thread algorithm based off Jamie Zawinski message threading.
//
// See: https://www.jwz.org/doc/threading.html
package database

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Linked list data structure for holding messages and place holder messages
type thread_container struct {
	message  *Message
	previous *thread_container
	next     *thread_container
}

func (container *thread_container) head() *thread_container {
	if container.previous == nil {
		return container
	}

	return container.previous.head()
}

func (container *thread_container) tail() *thread_container {
	if container.next == nil {
		return container
	}

	return container.next.tail()
}

func (container *thread_container) append(tc *thread_container) {
	if container.head() == tc.head() || tc.previous != nil {
		return
	}

	tail := container.tail()
	tc.previous = tail
	tail.next = tc
}

func normaliseSubject(subject string) string {
	subject = strings.TrimLeft(subject, "Re: ")

	return subject
}

func (container *thread_container) contains(tc *thread_container) bool {
	if container == tc || container.previous == tc || container.next == tc {
		return true
	}

	if container.next != nil {
		return container.next.contains(tc)
	}

	return false
}

func (container *thread_container) print(depth int) {
	if container.message != nil {
		fmt.Printf("%s\n", container.message.MessageID)
	}

	if container.next != nil {
		container.next.print(depth + 1)
	}
}

func (d *Database) UpdateThreads() error {
	// The id_table is a hash table associating Message-IDs with Containers.
	id_table := map[string]*thread_container{}

	it := d.messages.NewIter(nil)
	defer it.Close()
	// 1. For each message
	for it.First(); it.Valid(); it.Next() {
		message := Message{}
		if err := json.Unmarshal(it.Value(), &message); err != nil {
			return err
		}

		// A. If id_table contains an empty Container for this ID
		container, found := id_table[message.MessageID]
		if !found {
			// Create a new Container object holding this message
			container = &thread_container{message: &message}
			id_table[message.MessageID] = container
		} else if container.message == nil {
			// Store this message in the Container's message slot.
			container.message = &message
		}

		// For each element in the message's References field
		var reference_container *thread_container
		for _, reference := range message.References {
			// Find a Container object for the given Message-ID
			reference_container, found = id_table[reference]
			// If there's one in id_table use that; Otherwise, make (and index)
			// one with a null Message.
			if !found {
				reference_container = &thread_container{}
				id_table[reference] = reference_container
			}

			// Link the References field's Containers together in the order
			// implied by the References header
			container.head().append(reference_container)
		}

		// C. Set the parent of this message to be the last element in
		// References
		if container.next == nil && reference_container != nil && !reference_container.head().contains(container) {
			reference_container.append(container)
		}
	}

	// 2. ind the root set.
	thread_table := map[string]*thread_container{}
	for _, container := range id_table {
		if container.previous == nil {
			thread_table[container.message.MessageID] = container
		}
	}

	// 3. Discard id_table. We don't need it any more.
	id_table = nil

	// 4. Prune empty containers.
	for _, container := range id_table {
		for item := container; item != nil; item = item.next {
			if item.message == nil {
				prev := item.previous
				next := item.next

				prev.next = next
				next.previous = prev
			}
		}
	}

	// 5. Group root set by subject.
	// This not 100% how its is meant to be however this works OK for what
	// emails I have tested with. We can start to do other things like message
	// date as suggested
	subject_table := map[string]*thread_container{}
	for id, container := range thread_table {
		subject := normaliseSubject(container.message.Subject)
		thread, found := subject_table[subject]
		if !found {
			subject_table[subject] = container
			continue
		}

		thread.append(container)
		delete(thread_table, id)
	}

	// Nil out the subject_table so GC can collect it. Its not in the algorithm
	// however this should be the same as point 3
	subject_table = nil

	// 6. Now you're done threading!
	// Build and write the threads to the database
	for _, container := range thread_table {
		thread := Thread{
			ThreadID: d.NextThreadID(),
			Messages: []string{container.message.MessageID},
		}

		container.message.ThreadID = thread.ThreadID
		d.WriteMessage(*container.message)

		latest_message := container.message
		for item := container; item != nil; item = item.next {
			if item.message != nil && len(item.message.Account) > 0 {
				if item.message.Date.After(latest_message.Date) {
					latest_message = item.message
				}

				if !thread.HasMessage(item.message.MessageID) {
					thread.Messages = append(thread.Messages, item.message.MessageID)
				}

				for _, tag := range item.message.Tags {
					if !thread.HasTag(tag) {
						thread.Tags = append(thread.Tags, tag)
					}
				}
			}

			item.message.ThreadID = thread.ThreadID
			d.WriteMessage(*item.message)
		}

		thread.Date = latest_message.Date
		thread.Subject = latest_message.Subject
		thread.FromName = latest_message.FromName
		thread.From = latest_message.From

		d.WriteThread(thread)
	}

	return nil
}
