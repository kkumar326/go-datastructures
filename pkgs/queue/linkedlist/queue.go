package linkedlist

import (
	"errors"
)

/**
Queue Struct:
- Pointer to Head Element
- Pointer to Tail Element
- Size

Element Struct:
- Item of any type
- Pointer to previous Element address

Queue Methods
- NewQueue to create a new queue
- Enqueue to add Element to the tail of the queue
- Dequeue to remove Element from the head of the queue
- Len to find length
- IsEmpty to check if queue is empty
- Clear to clear queue
- Head returns head element
- Tail returns tail element
*/

type Queue struct {
	head *element
	tail *element
	size int
}

func NewQueue(item *ItemType) *Queue {
	if item != nil {
		element := &element{
			item:     *item,
			previous: nil,
		}

		return &Queue{
			head: element,
			tail: element,
			size: 1,
		}
	} else {
		return &Queue{
			head: nil,
			tail: nil,
			size: 0,
		}
	}
}

func (queue *Queue) Enqueue(item ItemType) ItemType {
	// new tail element
	newTail := &element{
		item:     item,
		previous: nil,
	}

	if queue.size == 0 {
		queue.head = newTail
		queue.tail = newTail
	} else {
		// set previous tail pointer to this new tail element
		queue.tail.previous = newTail

		// set new tail value
		queue.tail = newTail
	}

	// increase queue size
	queue.size++

	return item
}

func (queue *Queue) Dequeue() (ItemType, error) {
	if queue.size == 0 {
		return nil, errors.New("Queue is already empty")
	} else if queue.size == 1 {
		item := queue.head.item
		queue.Clear()
		return item, nil
	} else {
		item := queue.head.item

		// change head value to previous element
		queue.head = queue.head.previous
		queue.size--

		return item, nil
	}
}

func (queue *Queue) Len() int {
	return queue.size
}

func (queue *Queue) Head() (ItemType, error) {
	if queue.size != 0 {
		return queue.head.item, nil
	}

	return nil, errors.New("Queue is empty")
}

func (queue *Queue) Tail() (ItemType, error) {
	if queue.size != 0 {
		return queue.tail.item, nil
	}

	return nil, errors.New("Queue is empty")
}

func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *Queue) Clear() {
	queue.head = nil
	queue.tail = nil
	queue.size = 0
}
