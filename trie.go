package main

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	Children map[rune]*Node
}

func main() {
	root := &Node{}

	for _, fruit := range fruits {
		insertString(root, fruit)
	}

	words, err := findAutoCompletions(root, "Che")
	if err != nil {
		fmt.Println(err)
	} else {
		sort.Strings(words)
		fmt.Println(strings.Join(words, ", "))
	}

	// words := trieWords(root)
	// for _, w := range words {
	// 	fmt.Println(w)
	// }

	// yes := checkString(root, "Blood orange")
	// fmt.Printf("%v", yes)

	// fmt.Printf("digraph trie {\n")
	// dumpDot('_', 0, root)
	// fmt.Printf("}\n")
}

func insertString(root *Node, str string) {
	for _, v := range str {
		if root.Children == nil {
			root.Children = make(map[rune]*Node)
		}
		if root.Children[v] == nil {
			// insert node and advance
			root.Children[v] = &Node{
				Children: make(map[rune]*Node),
			}
		}

		root = root.Children[v]
	}
}

func checkString(root *Node, str string) bool {
	for _, v := range str {
		if root.Children[v] == nil {
			return false
		} else {
			root = root.Children[v]
		}
	}

	var childLen int
	for range root.Children {
		childLen++
	}

	return childLen == 0
}

func findAutoCompletions(root *Node, str string) ([]string, error) {
	for _, v := range str {
		root = root.Children[v]
		if root == nil {
			return nil, fmt.Errorf("%s word not found", str)
		}
	}
	words := trieWords(root, str)
	return words, nil
}

func trieWords(root *Node, str string) []string {
	var words []string
	var search func(node *Node, str string)

	search = func(node *Node, str string) {
		var childLen int
		for range node.Children {
			childLen++
		}
		if childLen == 0 {
			words = append(words, str)
		} else {
			for r, child := range node.Children {
				search(child, str+string(r))
			}
		}
	}

	search(root, str)
	return words
}

func dumpDot(rootc rune, i int, root *Node) {
	for c, child := range root.Children {
		fmt.Printf("    \"%d %c\" -> \"%d %c\";\n", i, rootc, i+1, c)
		dumpDot(c, i+1, child)
	}
}
