package trie

import (
	"unsafe"
)

type Trie struct {
	nodes   []node
	entries []nodeEntry
}

type node struct {
	isEnd bool
	next  []nodeEntry
}

type nodeEntry struct {
	chr byte
	idx uint16
}

func New() Trie {
	return Trie{
		nodes: []node{{
			isEnd: false,
			next:  nil,
		}},
	}
}

func fill(entries *[]nodeEntry, t *Trie, n *node) {
	*entries = append(*entries, n.next...)
	for i := range n.next {
		fill(entries, t, &t.nodes[n.next[i].idx])
	}
}

func pack(entries []nodeEntry, t *Trie, n *node, idx int) int {
	t.entries = entries
	n.next = t.entries[idx : idx+len(n.next)]
	idx += len(n.next)
	for i := range n.next {
		idx = pack(entries, t, &t.nodes[n.next[i].idx], idx)
	}
	return idx
}

func (t *Trie) Pack() int {
	entries := make([]nodeEntry, 0)
	fill(&entries, t, &t.nodes[0])
	pack(entries, t, &t.nodes[0], 0)
	t.entries = entries
	beg := unsafe.Pointer(&entries[0])
	end := unsafe.Pointer(&entries[1])
	return int(uintptr(end)-uintptr(beg)) * len(entries)
}

func (t *Trie) Add(str string) {
	n := 0
	bytes := *(*[]byte)(unsafe.Pointer(&str))
	for _, b := range bytes {
		nextNode := -1
		for j := range t.nodes[n].next {
			if t.nodes[n].next[j].chr == b {
				nextNode = int(t.nodes[n].next[j].idx)
				break
			}
		}
		if nextNode < 0 {
			idx := len(t.nodes)
			t.nodes = append(t.nodes, node{
				isEnd: false,
				next:  nil,
			})
			t.nodes[n].next = append(t.nodes[n].next, nodeEntry{chr: b, idx: uint16(idx)})
			nextNode = idx
		}
		n = nextNode
	}
	t.nodes[n].isEnd = true
}
func (t *Trie) Match(str string) bool {
	n := 0
	bytes := *(*[]byte)(unsafe.Pointer(&str))
	for _, b := range bytes {
		if t.nodes[n].isEnd {
			return true
		}
		nextNode := -1
		for j := range t.nodes[n].next {
			if t.nodes[n].next[j].chr == b {
				nextNode = int(t.nodes[n].next[j].idx)
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
