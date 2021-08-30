package stack

import "errors"

/**
Stack Struct:
- Pointer to top Element
- Size/ Length

Element Struct
- Item of any type
- Pointer to previous Element address

Stack methods:
NewStack to create a new stack
Len to find length
Push to push element in stack
Pop to pull element from stack
Peek to get top element
IsEmpty to check if stack is empty
Clear to clear stack

Important Point:
Mutex lock should be implemented across operations

Write Tests
*/

type Stack struct {
	top  *element
	size int
}

func NewStack(item *ItemType) *Stack {
	if item != nil {
		return &Stack{
			top: &element{
				item:     *item,
				previous: nil,
			},
			size: 1,
		}
	} else {
		return &Stack{
			top:  nil,
			size: 0,
		}
	}
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(item ItemType) ItemType {
	// changing stack top
	stack.top = &element{
		item:     item,
		previous: stack.top,
	}

	// incrementing stack size
	stack.size++

	return item
}

func (stack *Stack) Pop() (ItemType, error) {
	top := stack.top

	if stack.size >= 1 {
		// set top to the previous pointer of Top element
		stack.top = stack.top.previous

		// reduce stack size
		stack.size--

		return top.item, nil
	} else {
		return nil, errors.New("Stack is already empty")
	}
}

func (stack *Stack) Peek() (ItemType, error) {
	if stack.top != nil {
		return stack.top.item, nil
	} else {
		return nil, errors.New("Stack is empty")
	}
}

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) Clear() {
	stack.top = nil
	stack.size = 0
}
