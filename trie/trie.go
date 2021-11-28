package trie

import (
	"fmt"
)

type Node struct {
	Children map[rune]*Node
}

func NewTrie() *Node {
	return &Node{}
}

func (root *Node) InsertString(str string) {
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

func (root *Node) CheckString(str string) bool {
	for _, v := range str {
		if root.Children[v] == nil {
			return false
		} else {
			root = root.Children[v]
		}
	}

	for range root.Children {
		return false
	}

	return true
}

func (root *Node) FindAutoCompletions(str string) ([]string, error) {
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
			break
		}
		if childLen == 0 && len(str) > 0 {
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

func DumpDot(rootc rune, i int, root *Node) {
	for c, child := range root.Children {
		fmt.Printf("    \"%d %c\" -> \"%d %c\";\n", i, rootc, i+1, c)
		DumpDot(c, i+1, child)
	}
}
