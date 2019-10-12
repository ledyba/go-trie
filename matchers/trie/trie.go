package trie

import (
	"sort"
	"unsafe"
)

type Trie struct {
	nodes []node
}

func (n node) Len() int {
	return len(n)
}

func (n node) Less(i, j int) bool {
	return n[i].chr < n[j].chr
}

func (n node) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

type node []nodeEntry

type nodeEntry struct {
	chr     byte
	nextIdx uint16 // "0" means "isEnd". It's ok since trie is a tree, not a graph.
}

func New() *Trie {
	return &Trie{
		nodes: []node{{}},
	}
}

func FromWords(words []string) *Trie {
	tr := New()
	for _, str := range words {
		tr.Add(str)
	}
	tr.Pack()
	return tr
}

func (tr *Trie) Pack() int {
	return 0
}

func (tr *Trie) Add(str string) {
	n := 0
	bytes := *(*[]byte)(unsafe.Pointer(&str))
	lastIndex := len(bytes) - 1
	for ib, b := range bytes {
		nextNode := -1
		for j := range tr.nodes[n] {
			if tr.nodes[n][j].chr == b {
				if tr.nodes[n][j].nextIdx == 0 {
					return
				}
				nextNode = int(tr.nodes[n][j].nextIdx)
				break
			}
		}
		if nextNode < 0 {
			if ib == lastIndex { // last byte
				tr.nodes[n] = append(tr.nodes[n], nodeEntry{
					chr:     b,
					nextIdx: 0,
				})
			} else { // to be continued...
				idx := len(tr.nodes)
				tr.nodes = append(tr.nodes, node{})
				tr.nodes[n] = append(tr.nodes[n], nodeEntry{
					chr:     b,
					nextIdx: uint16(idx),
				})
				nextNode = idx
			}
			sort.Sort(tr.nodes[n])
		}
		n = nextNode
	}
}
func (tr *Trie) MatchBytesFrom(bytes []byte, from int) bool {
	nodes := tr.nodes
	if len(nodes[0]) == 0 {
		// Empty trie should match any string.
		return true
	}
	n := uint16(0)
	bytesLen := len(bytes)
	for ib := from; ib < bytesLen; ib++ {
		b := bytes[ib]
		nextNode := uint16(0)
		currentNode := nodes[n]
		currentNodeLen := len(currentNode)
		for i := 0; i < currentNodeLen; i++ {
			next := &currentNode[i]
			if b == next.chr {
				if next.nextIdx == 0 {
					return true
				}
				nextNode = next.nextIdx
				break
			} else if b < next.chr {
				break
			}
		}
		if nextNode == 0 {
			return false
		}
		n = nextNode
	}
	return false
}
func (tr *Trie) MatchBytes(bytes []byte) bool {
	return tr.MatchBytesFrom(bytes, 0)
}
func (tr *Trie) Match(str string) bool {
	return tr.MatchBytes(*(*[]byte)(unsafe.Pointer(&str)))
}
func (tr *Trie) Contains(str string) bool {
	nodes := tr.nodes
	if len(nodes[0]) == 0 {
		// Empty trie should match any string.
		return true
	}
	bytes := *(*[]byte)(unsafe.Pointer(&str))
	bytesLen := len(bytes)
	for i := 0; i < bytesLen; {
		n := uint16(0)
		bytesLen := len(bytes)
		for ib := i; ib < bytesLen; ib++ {
			b := bytes[ib]
			nextNode := uint16(0)
			currentNode := nodes[n]
			for _, next := range currentNode {
				if b == next.chr {
					if next.nextIdx == 0 {
						return true
					}
					nextNode = next.nextIdx
					break
				} else if b < next.chr {
					break
				}
			}
			if nextNode == 0 {
				break
			}
			n = nextNode
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
