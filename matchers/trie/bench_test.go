package trie

import (
	"testing"
)

func BenchmarkUnmatchRegex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TestRegex.MatchString(UnmatchTestString)
	}
}

func BenchmarkUnmatchTrie(b *testing.B) {
	tr, err := TestTrieTree(`../../words.txt`)
	if err != nil {
		b.Fatal(err)
	}
	tr.Pack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Contains(UnmatchTestString)
	}
}
