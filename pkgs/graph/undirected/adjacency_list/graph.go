package adjacencylist

import (
	"errors"
)

type graph struct {
	nodes []*Node
}

func NewGraph() *graph {
	return &graph{}
}

func (gph *graph) addNode(node *Node) {
	gph.nodes = append(gph.nodes, node)
}

func (gph *graph) addEdge(source, dest *Node) error {
	var validSource, validDest = false, false

	// check valid source and destination nodes
	// nodes should be in node list
	for _, node := range gph.nodes {
		if validSource && validDest {
			break
		}

		if node == source {
			validSource = true
		}

		if node == dest {
			validDest = true
		}
	}

	if !validSource {
		return errors.New("invalid source node")
	}

	if !validDest {
		return errors.New("invalid destination node")
	}

	currentNode := source
	for {
		if currentNode.next == nil {
			currentNode.next = dest
			break
		}

		currentNode = currentNode.next
	}

	return nil
}

func (gph *graph) BFS(start *Node) {

}

func (gph *graph) DFS(start *Node) {

}
