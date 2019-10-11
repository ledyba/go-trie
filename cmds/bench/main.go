package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ledyba/go-trie/matchers/test_util"
	"github.com/ledyba/go-trie/matchers/trie"
)

func main() {
	words, err := test_util.ReadWords("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	tr := trie.FromWords(words)
	var regexTime float64
	var trieTime float64
	if false {
		N := 1000
		n := 0
		beg := time.Now()
		for i := 0; i < N; i++ {
			if !test_util.TestRegex.MatchString(test_util.UnmatchTestString) {
				n++
			}
		}
		end := time.Now()
		elapsedInMicro := end.Sub(beg).Microseconds()
		regexTime = float64(elapsedInMicro) / float64(n)
		fmt.Printf(" - golang's trie2regex: %.2f us / op (%d times)\n", regexTime, n)
	}

	{
		N := 100000
		n := 0
		beg := time.Now()
		for i := 0; i < N; i++ {
			if !tr.Contains(test_util.UnmatchTestString) {
				n++
			}
		}
		end := time.Now()
		elapsedInMicro := end.Sub(beg).Microseconds()
		trieTime = float64(elapsedInMicro) / float64(n)
		fmt.Printf(" - go-trie: %.2f us/op (x %.2f) (%d times)\n", trieTime, regexTime/trieTime, n)
	}
}
