package leetcode

import "fmt"

/*
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.*/

// 字典树,实现还算简单,有时间分解一下

type Trie struct {
	End bool
	Ch  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var trie Trie
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		idx := c - 'a'
		if cur.Ch[idx] == nil {
			var newnode Trie
			cur.Ch[idx] = &newnode
			cur = &newnode
		} else {
			cur = cur.Ch[idx]
		}
	}
	cur.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range word {
		idx := c - 'a'
		fmt.Printf("%c,", c)
		if cur.Ch[idx] == nil {
			return false
		} else {
			cur = cur.Ch[idx]
		}
	}
	return cur.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		idx := c - 'a'
		if cur.Ch[idx] == nil {
			return false
		} else {
			cur = cur.Ch[idx]
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
