package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/malisetti/prefix-tree/trie"
)

func main() {
	root := trie.NewTrie()

	for _, fruit := range fruits {
		root.Insert(fruit)
	}

	words, err := root.Completions("Che")
	if err != nil {
		fmt.Println(err)
	} else {
		sort.Strings(words)
		fmt.Println(strings.Join(words, ", "))
	}

	isWord, isSubStr := root.Check("Cherry")
	fmt.Printf("%s is word %v, is substr %v\n", "Cherry", isWord, isSubStr)

	isWord, isSubStr = root.Check("Blood")
	fmt.Printf("%s is word %v, is substr %v\n", "Blood", isWord, isSubStr)

	fmt.Printf("digraph trie {\n")
	trie.DumpDot('_', 0, root)
	fmt.Printf("}\n")
}
