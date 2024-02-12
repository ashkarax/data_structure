package main

import "fmt"

const size = 26

type Node struct {
	children [size]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func (trie *Trie) Insert(str string) {
	n := len(str)
	current := trie.root

	for i := 0; i < n; i++ {
		charIdx := str[i] - 'a'

		if current.children[charIdx] == nil {
			current.children[charIdx] = &Node{}
		}
		current = current.children[charIdx]
	}
	current.isEnd = true
}

func (trie *Trie) Search(str string) bool {
	n := len(str)
	current := trie.root

	for i := 0; i < n; i++ {
		charIdx := str[i] - 'a'

		if current.children[charIdx] == nil {
			return false
		}
		fmt.Println(current.children[charIdx])
		current = current.children[charIdx]
	}
	return current.isEnd
}

func main() {
	trie := &Trie{root: &Node{}}

	var v = []string{
		"hello",
		"banana",
		"bamboo",
	}

	for _, s := range v {
		trie.Insert(s)
	}

	fmt.Println(trie.Search("hello"))
}
