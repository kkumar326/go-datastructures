package adjacencylist

import (
	"errors"

	graphDS "github.com/kkumar326/go-dsa/pkgs/graph"
)

type graph struct {
	nodes []*graphDS.Node
}

func NewGraph() *graph {
	return &graph{}
}

func (gph *graph) addNode(node *graphDS.Node) {
	gph.nodes = append(gph.nodes, node)
}

func (gph *graph) addEdge(from, to *graphDS.Node) error {
	var validFrom, validTo = false, false

	// check valid from and to nodes
	// nodes should be in node list
	for _, node := range gph.nodes {
		if validFrom && validTo {
			break
		}

		if node == from {
			validFrom = true
		}

		if node == to {
			validTo = true
		}
	}

	if !validFrom {
		return errors.New("invalid from node")
	}

	if !validTo {
		return errors.New("invalid to node")
	}

	from.Edges = append(from.Edges, to)

	return nil
}
