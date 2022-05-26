package trie

// ref: http://www.code2succeed.com/golang-insert-and-search-trie/
// https://golangbyexample.com/trie-implementation-in-go/

func New() *trie {
	return &trie{
		root: &Node{},
	}
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &Node{}
		}
		current = current.children[index]
	}
	current.isWordEnd = true
}

func (t *trie) search(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current.isWordEnd
}

func (t *trie) delete(word string) bool {
	return true
}
