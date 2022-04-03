package graph

type ItemType interface{}

type Node struct {
	Value ItemType
	Edges []*Node
}
