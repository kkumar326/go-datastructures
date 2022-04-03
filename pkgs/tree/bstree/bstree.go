package bstree

type bstree interface {
	insert(key int, value ItemType)
	search(key int)
	inOrderTraverse(node *Node)
	preOrderTraverse(node *Node)
	postOrderTraverse(node *Node)
	levelOrderTraverse(node *Node)
	min() int
	max() int
	remove(key int) *Node
	stringify()
}
