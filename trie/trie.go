package trie

import (
	"fmt"
)

type Trie struct {
	root *node
}

type node struct {
	children map[rune]*node
}

func NewTrie() *Trie {
	return &Trie{
		root: &node{},
	}
}

func (trie *Trie) Insert(str string) {
	root := trie.root
	for _, v := range str {
		if root.children == nil {
			root.children = make(map[rune]*node)
		}
		if root.children[v] == nil {
			// insert node and advance
			root.children[v] = &node{
				children: make(map[rune]*node),
			}
		}

		root = root.children[v]
	}
}

func (trie *Trie) Check(str string) (isWord, isSubStr bool) {
	root := trie.root
	for _, v := range str {
		if root.children[v] == nil {
			return
		} else {
			root = root.children[v]
		}
	}

	for range root.children {
		return false, true
	}

	return true, false
}

func (trie *Trie) Completions(str string) ([]string, error) {
	root := trie.root
	for _, v := range str {
		root = root.children[v]
		if root == nil {
			return nil, fmt.Errorf("%s word not found", str)
		}
	}
	words := words(root, str)
	return words, nil
}

func words(root *node, str string) []string {
	var words []string
	var search func(node *node, str string)

	search = func(node *node, str string) {
		if len(root.children) == 0 && len(str) > 0 {
			words = append(words, str)
		} else {
			for r, child := range node.children {
				search(child, str+string(r))
			}
		}
	}

	search(root, str)
	return words
}

func DumpDot(rootc rune, i int, trie *Trie) {
	var dump func(rootc rune, i int, root *node)
	dump = func(rootc rune, i int, root *node) {
		for c, child := range root.children {
			fmt.Printf("    \"%d %c\" -> \"%d %c\";\n", i, rootc, i+1, c)
			dump(rune(c), i+1, child)
		}
	}
	dump(rootc, i, trie.root)
}
