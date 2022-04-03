package linkedlist

import (
	"fmt"

	bsTree "github.com/kkumar326/go-dsa/pkgs/tree/bstree"
)

type bstree struct {
	root *bsTree.Node
}

func New(key int, value bsTree.ItemType) *bstree {
	return &bstree{
		root: &bsTree.Node{
			Key:   key,
			Value: value,
		},
	}
}

func (bst *bstree) insert(key int, value bsTree.ItemType) {
	newElement := bsTree.Node{
		Key:   key,
		Value: value,
	}

	if bst.root == nil {
		bst.root = &newElement
	} else {
		insertNode(bst.root, &newElement)
	}
}

func insertNode(node *bsTree.Node, newElement *bsTree.Node) {
	if newElement.Key < node.Key {
		if node.LeftChild == nil {
			node.LeftChild = newElement
		} else {
			insertNode(node.LeftChild, newElement)
		}
	} else {
		if node.RightChild == nil {
			node.RightChild = newElement
		} else {
			insertNode(node.RightChild, newElement)
		}
	}
}

func (bst *bstree) search(key int) bool {
	return search(bst.root, key)
}

func search(node *bsTree.Node, key int) bool {
	if key == node.Key {
		return true
	}

	if key < node.Key {
		if node.LeftChild != nil {
			search(node.LeftChild, key)
		}
	} else {
		if node.RightChild != nil {
			search(node.RightChild, key)
		}
	}

	return false
}

func (bst *bstree) inOrderTraverse(node *bsTree.Node) {
	inOrderTraverse(bst.root)
}

func inOrderTraverse(node *bsTree.Node) {
	if node == nil {
		return
	}

	inOrderTraverse(node.LeftChild)

	fmt.Printf("%v->", node.Key)

	inOrderTraverse(node.RightChild)
}

func (bst *bstree) preOrderTraverse(node *bsTree.Node) {
	preOrderTraverse(bst.root)
}

func preOrderTraverse(node *bsTree.Node) {
	if node == nil {
		return
	}

	fmt.Printf("%v->", node.Key)

	preOrderTraverse(node.LeftChild)
	preOrderTraverse(node.RightChild)
}

func (bst *bstree) postOrderTraverse(node *bsTree.Node) {
	postOrderTraverse(bst.root)
}

func postOrderTraverse(node *bsTree.Node) {
	if node == nil {
		return
	}

	preOrderTraverse(node.LeftChild)
	preOrderTraverse(node.RightChild)

	fmt.Printf("%v->", node.Key)
}

func levelOrderTraverse(node *bsTree.Node) {

}

func (bst *bstree) min() int {
	return min(bst.root)
}

func min(node *bsTree.Node) int {
	minKey := node.Key

	for {
		if node.LeftChild != nil {
			node = node.LeftChild
		} else {
			minKey = node.Key
			break
		}
	}

	return minKey
}

func (bst *bstree) max() int {
	return max(bst.root)
}

func max(node *bsTree.Node) int {
	maxKey := node.Key

	for {
		if node.RightChild != nil {
			node = node.RightChild
		} else {
			maxKey = node.Key
			break
		}
	}

	return maxKey
}

func (bst *bstree) remove(key int) {
	remove(bst.root, key)
}

// ref: https://gist.github.com/mycodeschool/9465a188248b624afdbf
func remove(node *bsTree.Node, key int) *bsTree.Node {
	if node == nil {
		return nil
	} else if key > node.Key {
		remove(node.RightChild, key)
	} else if key < node.Key {
		remove(node.LeftChild, key)
	} else {
		// no child
		if node.LeftChild == nil && node.RightChild == nil {
			temp := node
			node = nil
			return temp
		}

		// one child
		if node.LeftChild == nil && node.RightChild != nil {
			return nil
		}

		// both children
		return nil
	}
	return nil
}

func (bst *bstree) stringify() {}
