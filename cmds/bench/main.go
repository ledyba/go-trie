package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ledyba/go-trie/types/trie"
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
		fmt.Printf(" - golang's trie2regexp (%d times): %.2f us / op\n", N, regexTime)
	}

	{
		N := 10000
		beg := time.Now()
		for i := 0; i < N; i++ {
			tr.Contains(trie.UnmatchTestString)
		}
		end := time.Now()
		elapsedInMicro := end.Sub(beg).Microseconds()
		trieTime = float64(elapsedInMicro) / float64(N)
		fmt.Printf(" - go-trie (%d times): %.2f us / op (x %.2f)\n", N, trieTime, regexTime/trieTime)
	}
}
