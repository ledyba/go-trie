package main

import (
	"fmt"
	"time"

	"github.com/ledyba/go-trie/matchers/test_util"
	"github.com/ledyba/go-trie/matchers/trie"
)

func main() {
	tr := trie.FromWords(test_util.Words)
	N := 1000000
	n := 0
	beg := time.Now()
	for i := 0; i < N; i++ {
		if !tr.Contains(test_util.UnmatchString) {
			n++
		}
	}
	end := time.Now()
	elapsedInMicro := end.Sub(beg).Microseconds()
	fmt.Printf(" - go-trie: %.2f us/op (%d times)\n", float64(elapsedInMicro)/float64(n), n)
}
