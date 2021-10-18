package bstree

type bstree interface {
	insert(item interface{})
	search(item interface{})
	inOrderTraverse()
	preOrderTraverse()
	postOrderTraverse()
	min()
	max()
	remove(item interface{})
	stringify()
}
