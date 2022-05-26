package trie

// we work with english alphabets only
// reason: https://stackoverflow.com/a/10229225
const ALBHABET_SIZE = 26

type Node struct {
	children  [ALBHABET_SIZE]*Node
	isWordEnd bool
}

type trie struct {
	root *Node
}
