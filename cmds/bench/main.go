package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ledyba/go-trie/matchers/trie"
)

func main() {
	tr, err := trie.TestTrieTree("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	var regexTime float64
	var trieTime float64
	if false {
		N := 1000
		beg := time.Now()
		for i := 0; i < N; i++ {
			trie.TestRegex.MatchString(trie.UnmatchTestString)
		}
		end := time.Now()
		elapsedInMicro := end.Sub(beg).Microseconds()
		regexTime = float64(elapsedInMicro) / float64(N)
		fmt.Printf(" - golang's trie2regex: %.2f us / op (%d times)\n", regexTime, N)
	}

	{
		N := 100000
		beg := time.Now()
		for i := 0; i < N; i++ {
			tr.Contains(trie.UnmatchTestString)
		}
		end := time.Now()
		elapsedInMicro := end.Sub(beg).Microseconds()
		trieTime = float64(elapsedInMicro) / float64(N)
		fmt.Printf(" - go-trie: %.2f us/op (x %.2f) (%d times)\n", trieTime, regexTime/trieTime, N)
	}
}
