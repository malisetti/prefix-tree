package main

import (
	"fmt"
)

type Node struct {
	Children map[rune]*Node
}

func main() {
	root := &Node{}

	// for _, fruit := range fruits {
	// 	insertString(root, fruit)
	// }

	insertString(root, "hello")
	insertString(root, "helium")

	// words, err := findAutoCompletions(root, "hel")
	// fmt.Println(words, err)

	displayWords(root, "")

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

	return true
}

func findAutoCompletions(root *Node, str string) ([]string, error) {
	for _, v := range str {
		root = root.Children[v]
		if root == nil {
			return nil, fmt.Errorf("%s word not found", str)
		}
	}

	return nil, nil
}

func displayWords(root *Node, str string) {
	for r, child := range root.Children {
		strx := str + string(r)

		displayWords(child, strx)
	}

	println(str)
}

func dumpDot(rootc rune, i int, root *Node) {
	for c, child := range root.Children {
		fmt.Printf("    \"%d %c\" -> \"%d %c\";\n", i, rootc, i+1, c)
		dumpDot(c, i+1, child)
	}
}
