package linkedlist

import (
	bsTree "github.com/kkumar326/go-dsa/pkgs/tree/bstree"
)

/*
Traversal pseudo code:

In Order
- Start on root node
- If node has left child not nil
-- Go to left child node and repeat
- Else
-- Fetch the node value
-- Check the right node
-- If right node is not nil
--- If node has left child not nil
....... same repetition
--- Else
....... same repetition
-- Else
--- Go back to parent
---- Fetch the node value
---- Check the right node
......... same repetition

func traverse(node) {
	if (node->left != nil) {
		traverse(node->left)
	} else {
		fmt.Println("%v", node->item)
		if (node->right != nil) {
			if (node->right->left != nil) {
				traverse(node)
			} else {
				fmt. Println()
				// repetition
			}
		} else {
			fmt.Println("%v", node->parent)

			if (node->parent->right != nil) {
				// repetition
			} else {
				// repetition
			}
		}
	}
}

*/

type bstree struct {
	item       bsTree.ItemType
	leftChild  *bstree
	rightChild *bstree
}
