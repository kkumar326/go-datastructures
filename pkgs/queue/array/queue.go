package array

import (
	"errors"

	q "github.com/kkumar326/go-dsa/pkgs/queue"
)

// Ref: https://www.cs.usfca.edu/~galles/visualization/QueueArray.html

// TODO: overflow checks

type Queue struct {
	queueList []q.ItemType
	head      int // array index
	tail      int // array index
	size      int // filled spaces, not capacity
}

func NewQueue(item *q.ItemType) *Queue {
	if item != nil {
		return &Queue{
			queueList: []q.ItemType{*item},
			head:      0,
			tail:      0,
			size:      1,
		}
	}

	return &Queue{}
}

func (queue *Queue) Enqueue(item q.ItemType) q.ItemType {
	if queue.size == 0 {
		queue.queueList[0] = item
	} else {
		// add item to slice
		queue.queueList[queue.tail+1] = item

		// move tail to next index
		queue.tail++
	}

	// increase queue size
	queue.size++

	return item
}

func (queue *Queue) Dequeue() (q.ItemType, error) {
	if queue.size == 0 {
		return nil, errors.New("Queue is already empty")
	} else if queue.size == 1 {
		item := queue.queueList[queue.head]
		queue.Clear()
		return item, nil
	} else {
		item := queue.queueList[queue.head]

		// unset the head value to nil
		queue.queueList[queue.head] = nil

		// move head index to next one
		queue.head++

		// reduce size by "1"
		queue.size--

		return item, nil
	}
}

func (queue *Queue) Len() int {
	return queue.size
}

func (queue *Queue) Head() (q.ItemType, error) {
	if queue.size == 0 {
		return nil, errors.New("Queue is empty")
	}

	return queue.queueList[queue.head], nil
}

func (queue *Queue) Tail() (q.ItemType, error) {
	if queue.size == 0 {
		return nil, errors.New("Queue is empty")
	}

	return queue.queueList[queue.tail], nil
}

func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *Queue) Clear() {
	queue.queueList = []q.ItemType{}
	queue.head = 0
	queue.tail = 0
	queue.size = 0
}
