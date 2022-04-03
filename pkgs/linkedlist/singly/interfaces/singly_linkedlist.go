package interfaces

import "github.com/kkumar326/go-dsa/pkgs/linkedlist/singly"

type SinglyLinkedList interface {
	Insert(item singly.ItemType)
	InsertAt()
	PrintAll()
	GetByVal()
	GetAt()
	DeleteAt()
	DeleteByVal()
}
