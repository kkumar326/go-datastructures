package bstree

type ItemType interface{}

type Node struct {
	Key        int
	Value      ItemType
	LeftChild  *Node
	RightChild *Node
}
