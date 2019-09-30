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
	{
		N := 10000
		fmt.Printf("[Baseline] benchmarking golang's regexp (%d times).\n", N)
		beg := time.Now()
		for i := 0; i < N; i++ {
			trie.TestRegex.MatchString(trie.UnmatchTestString)
		}
		end := time.Now()
		elapsedInMicro := end.Sub(beg).Microseconds()
		fmt.Printf(" - result: %f us / op\n", float64(elapsedInMicro)/float64(N))
		regexTime = float64(elapsedInMicro) / float64(N)
	}

	fmt.Println()

	{
		N := 10000
		fmt.Printf("[Ours] benchmarking trie (%d times).\n", N)
		beg := time.Now()
		for i := 0; i < N; i++ {
			tr.Contains(trie.UnmatchTestString)
		}
		end := time.Now()
		elapsedInMicro := end.Sub(beg).Microseconds()
		fmt.Printf(" - result: %f us / op\n", float64(elapsedInMicro)/float64(N))
		trieTime = float64(elapsedInMicro) / float64(N)
	}

	log.Printf("Trie is faster %.2f fastar than regex!", regexTime/trieTime)
}
