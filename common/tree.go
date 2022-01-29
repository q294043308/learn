package common

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 208. Implement Trie (Prefix Tree)
type Trie struct {
	key   []*Trie
	exist bool
}

func Constructor() Trie {
	return Trie{
		key: make([]*Trie, 256),
	}
}

func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.key[v] == nil {
			this.key[v] = &Trie{key: make([]*Trie, 256)}
		}

		this = this.key[v]
	}
	this.exist = true
}

func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.key[v] == nil {
			return false
		}

		this = this.key[v]
	}
	return this.exist
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.key[v] == nil {
			return false
		}

		this = this.key[v]
	}
	return true
}
