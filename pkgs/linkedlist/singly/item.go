package singly

type ItemType interface{}

type element struct {
	item ItemType
	next *element
}
