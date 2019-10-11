package trie

import (
	"testing"

	derekparker "github.com/derekparker/trie"
	"github.com/ledyba/go-trie/matchers/test_util"
)

func BenchmarkUnmatchDghubbleTrie(b *testing.B) {
	tr := derekparker.New()
	words, err := test_util.ReadWords(`../../words.txt`)
	if err != nil {
		b.Fatal(err)
	}
	for _, str := range words {
		tr.Add(str, struct{}{})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range test_util.UnmatchTestString {
			if tr.HasKeysWithPrefix(test_util.UnmatchTestString[j:]) {
				b.Error("Should Unmatch!")
			}
		}
	}
}

func BenchmarkUnmatchRegex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if test_util.TestRegex.MatchString(test_util.UnmatchTestString) {
			b.Error("Should Unmatch!")
		}
	}
}
func BenchmarkUnmatchTrie(b *testing.B) {
	words, err := test_util.ReadWords(`../../words.txt`)
	if err != nil {
		b.Fatal(err)
	}
	tr := FromWords(words)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if tr.Contains(test_util.UnmatchTestString) {
			b.Error("Should Unmatch!")
		}
	}
}
