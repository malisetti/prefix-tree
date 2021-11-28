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
		root.InsertString(fruit)
	}

	words, err := root.FindAutoCompletions("Che")
	if err != nil {
		fmt.Println(err)
	} else {
		sort.Strings(words)
		fmt.Println(strings.Join(words, ", "))
	}

	yes := root.CheckString("Blood orange")
	fmt.Printf("%v\n", yes)

	fmt.Printf("digraph trie {\n")
	trie.DumpDot('_', 0, root)
	fmt.Printf("}\n")
}
