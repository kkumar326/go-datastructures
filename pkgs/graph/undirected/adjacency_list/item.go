package adjacencylist

type ItemType interface{}

type Node struct {
	value  ItemType
	weight int // edge weight from source to destination
	next   *Node
}
