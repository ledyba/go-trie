package trie

type Trie struct {
	nodes []node
}

type node struct {
	isEnd bool
	next []struct{
		chr byte
		idx int
	}
}

func New() Trie {
	return Trie{
		nodes: []node{{
			isEnd: false,
			next:  nil,
		}},
	}
}

func (t *Trie) Add(str string) {
	n := 0
	for i := 0; i < len(str); i++ {
		s := str[i]
		nextNode := -1
		for j := range t.nodes[n].next {
			if t.nodes[n].next[j].chr == s{
				nextNode = t.nodes[n].next[j].idx
				break
			}
		}
		if nextNode < 0 {
			idx := len(t.nodes)
			t.nodes = append(t.nodes, node{
				isEnd: false,
				next:  nil,
			})
			t.nodes[n].next = append(t.nodes[n].next, struct {
				chr  byte
				idx int
			}{chr: s, idx: idx})
			nextNode = idx
		}
		n = nextNode
	}
	t.nodes[n].isEnd = true
}
func (t *Trie) Match(str string) bool {
	n := 0
	for i := 0; i < len(str); i++ {
		if t.nodes[n].isEnd {
			return true
		}
		nextNode := -1
		s := str[i]
		for j := range t.nodes[n].next {
			if t.nodes[n].next[j].chr == s {
				nextNode = t.nodes[n].next[j].idx
				break
			}
		}
		if nextNode < 0 {
			return false
		}
		n = nextNode
	}
	return t.nodes[n].isEnd
}
func (t *Trie) Contains(str string) bool {
	for i := 0; i < len(str); i++ {
		match := t.Match(str[i:])
		if match {
			return true
		}
	}
	return false
}
