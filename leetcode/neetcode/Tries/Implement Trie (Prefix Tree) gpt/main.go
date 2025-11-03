package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// データ構造
// (root)
//  └── a
//       └── p
//            ├── p (isEnd=true)
//            │    └── l
//            │         └── e (isEnd=true)
//            └── e (isEnd=true)

func (this *Trie) Insert(word string) {
	node := this.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// データ構造
// (root)
//  └── a
//       └── p
//            ├── p (isEnd=true)
//            │    └── l
//            │         └── e (isEnd=true)
//            └── e (isEnd=true)

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		// childrenはmapでできえてるんだ
		// map構造がネストしてる
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return true
}

func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true ("apple" は存在する)
	fmt.Println(trie.Search("app"))     // false ("app" はまだ登録されていない)
	fmt.Println(trie.StartsWith("app")) // true ("apple" が "app" で始まる)

	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true (今度は "app" も登録した)
}
