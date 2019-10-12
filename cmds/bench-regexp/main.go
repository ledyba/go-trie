package main

import (
	"fmt"
	"time"

	"github.com/ledyba/go-trie/matchers/test_util"
)

func main() {
	N := 1000
	n := 0
	beg := time.Now()
	for i := 0; i < N; i++ {
		if !test_util.RegexPattern.MatchString(test_util.UnmatchString) {
			n++
		}
	}
	end := time.Now()
	elapsedInMicro := end.Sub(beg).Microseconds()
	fmt.Printf(" - golang's trie2regex: %.2f us/op (%d times)\n", float64(elapsedInMicro)/float64(n), n)
}
