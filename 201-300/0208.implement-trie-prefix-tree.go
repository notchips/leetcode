/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */
 type Trie struct {
	end   bool
	edges []*edge
}

type edge struct {
	char byte
	next *Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		end:   false,
		edges: nil,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.end = true
		return
	}
	for _, e := range this.edges {
		if e.char == word[0] {
			e.next.Insert(word[1:])
			return
		}
	}
	e := &edge{
		char: word[0],
		next: new(Trie),
	}
	this.edges = append(this.edges, e)
	e.next.Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.end
	}
	for _, e := range this.edges {
		if e.char == word[0] {
			return e.next.Search(word[1:])
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	for _, e := range this.edges {
		if e.char == prefix[0] {
			return e.next.StartsWith(prefix[1:])
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
 