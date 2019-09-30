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
	t, err := TestTrieTree(`../../words.txt`)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Contains(UnmatchTestString)
	}
}
