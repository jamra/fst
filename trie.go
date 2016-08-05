package fst

//TrieNode is the node part of a trie
type TrieNode struct {
	letter      rune
	transitions map[rune]*TrieNode
}

func newTrieNode(letter rune) *TrieNode {
	t := &TrieNode{letter: letter, transitions: make(map[byte]*TrieNode)}
	return t
}


func (t *TrieNode) FindPrefix(prefix string) (*TrieNode, bool) {
    cur := t

    for letter := range prefix {
        if ok, t := cur.transitions[letter]; ok {
            cur = t
        } else {
            return cur, false
        }
    }
    return cur, false
}