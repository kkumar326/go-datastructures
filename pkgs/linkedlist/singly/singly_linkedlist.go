package singly

import "fmt"

type SinglyLinkedList struct {
	head *element
	tail *element
	size int
}

func NewSinglyLinkedList(item *ItemType) *SinglyLinkedList {
	if item != nil {
		element := &element{
			item: item,
			next: nil,
		}

		return &SinglyLinkedList{
			head: element,
			tail: element,
			size: 1,
		}
	} else {
		return &SinglyLinkedList{
			head: nil,
			tail: nil,
			size: 0,
		}
	}
}

func (s *SinglyLinkedList) Insert(item ItemType) {
	newElement := element{
		item: item,
		next: nil,
	}

	if s.size == 0 {
		s.head = &newElement
		s.tail = &newElement
		s.size++
	} else if s.size == 1 {
		s.head.next = &newElement
		s.tail = &newElement
		s.size++
	} else {
		s.tail.next = &newElement
		s.tail = &newElement
	}
}

func (s *SinglyLinkedList) InsertAt(item ItemType, index int) error {
	if index < 0 || index > s.size {
		return fmt.Errorf("index out of bounds")
	}

	nodeAtPrevIndex := s.head

	// running loop to iterate to 'index' to find existing node
	for i := 0; i < index-1; i++ {
		nodeAtPrevIndex = nodeAtPrevIndex.next
	}

	newElement := element{
		item: item,
		next: nodeAtPrevIndex.next,
	}

	nodeAtPrevIndex.next = &newElement
	s.size++

	return nil
}

func (s *SinglyLinkedList) PrintAll() {
	currentPointer := s.head

	for i := 0; i < s.size; i++ {
		fmt.Printf("Index: %d, Item: %v, Next Node: %v\n", i, currentPointer.item, currentPointer.next)
		currentPointer = currentPointer.next
	}
}

func (s *SinglyLinkedList) GetByVal(item ItemType) int {
	currentPointer := s.head

	for i := 0; i < s.size; i++ {
		if currentPointer.item == item {
			return i
		}
		currentPointer = currentPointer.next
	}

	return -1
}

func (s *SinglyLinkedList) GetAt(index int) (element, error) {
	if index < 0 || index > s.size {
		return element{}, fmt.Errorf("index out of bounds")
	}

	nodeAtIndex := s.head

	// running loop to iterate to 'index' to find existing node
	for i := 0; i < index; i++ {
		nodeAtIndex = nodeAtIndex.next
	}

	return *nodeAtIndex, nil
}

func (s *SinglyLinkedList) DeleteAt(index int) (element, error) {
	if index < 0 || index > s.size {
		return element{}, fmt.Errorf("index out of bounds")
	}

	nodeAtIndex := s.head

	// running loop to iterate to 'index' to find existing node
	for i := 0; i < index-1; i++ {
		nodeAtIndex = nodeAtIndex.next
	}

	nodeAtIndex.next = nodeAtIndex.next.next

	return *nodeAtIndex.next, nil
}

func (s *SinglyLinkedList) DeleteByVal(item ItemType) (element, error) {
	currentPointer := s.head
	previousPointer := s.head

	for i := 0; i < s.size; i++ {
		if currentPointer.item == item {
			previousPointer.next = previousPointer.next.next
			return *currentPointer, nil
		}

		if i > 1 {
			previousPointer = previousPointer.next
		}

		currentPointer = currentPointer.next
	}

	return element{}, fmt.Errorf("item not found")
}
