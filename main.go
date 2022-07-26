package main

import (
	"fmt"

	prefixtree "github.com/malisetti/prefix-tree/trie"
)

func main() {
	trie := prefixtree.NewTrie()

	for _, fruit := range fruits {
		trie.Insert(fruit)
	}

	fmt.Printf("digraph trie {\n")
	prefixtree.DumpDot('_', 0, trie)
	fmt.Printf("}\n")
}
