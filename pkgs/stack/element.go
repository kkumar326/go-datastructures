package stack

type ItemType interface{}

type element struct {
	item     ItemType
	previous *element
}
