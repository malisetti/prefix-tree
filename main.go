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

	words, err = root.Completions("Bla")
	if err != nil {
		fmt.Println(err)
	} else {
		sort.Strings(words)
		fmt.Println(strings.Join(words, ", "))
	}

	isWord, isSubStr := root.Check("Blood")
	fmt.Printf("%v %v\n", isWord, isSubStr)

	fmt.Printf("digraph trie {\n")
	trie.DumpDot('_', 0, root)
	fmt.Printf("}\n")
}
