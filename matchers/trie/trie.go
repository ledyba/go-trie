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

func New() *Trie {
	return &Trie{
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

func (tr *Trie) Pack() int {
	entries := make([]nodeEntry, 0)
	fill(&entries, tr, &tr.nodes[0])
	pack(entries, tr, &tr.nodes[0], 0)
	tr.entries = entries
	beg := unsafe.Pointer(&entries[0])
	end := unsafe.Pointer(&entries[1])
	return int(uintptr(end)-uintptr(beg)) * len(entries)
}

func (tr *Trie) Add(str string) {
	n := 0
	bytes := *(*[]byte)(unsafe.Pointer(&str))
	for _, b := range bytes {
		nextNode := -1
		for j := range tr.nodes[n].next {
			if tr.nodes[n].next[j].chr == b {
				nextNode = int(tr.nodes[n].next[j].idx)
				break
			}
		}
		if nextNode < 0 {
			idx := len(tr.nodes)
			tr.nodes = append(tr.nodes, node{
				isEnd: false,
				next:  nil,
			})
			tr.nodes[n].next = append(tr.nodes[n].next, nodeEntry{chr: b, idx: uint16(idx)})
			nextNode = idx
		}
		n = nextNode
	}
	tr.nodes[n].isEnd = true
}
func (tr *Trie) MatchBytesFrom(bytes []byte, from int) bool {
	n := &tr.nodes[0]
	for ib := from; ib < len(bytes); ib++ {
		b := bytes[ib]
		if n.isEnd {
			return true
		}
		var nextNode *node = nil
		nexts := n.next
		for j := range nexts {
			next := &n.next[j]
			if next.chr == b {
				nextNode = &tr.nodes[next.idx]
				break
			}
		}
		if nextNode == nil {
			return false
		}
		n = nextNode
	}
	return n.isEnd
}
func (tr *Trie) MatchBytes(bytes []byte) bool {
	return tr.MatchBytesFrom(bytes, 0)
}
func (tr *Trie) Match(str string) bool {
	return tr.MatchBytes(*(*[]byte)(unsafe.Pointer(&str)))
}
func (tr *Trie) Contains(str string) bool {
	bytes := *(*[]byte)(unsafe.Pointer(&str))
	bytesLen := len(bytes)
	for i := 0; i < bytesLen; {
		if tr.MatchBytesFrom(bytes, i) {
			return true
		}
		// see: https://tools.ietf.org/html/rfc3629
		b := bytes[i]
		if b&0b11110000 == 0b11110000 {
			i += 4
		} else if b&0b11100000 == 0b11100000 {
			i += 3
		} else if b&0b11000000 == 0b11000000 {
			i += 2
		} else {
			i += 1
		}
	}
	return false
}
