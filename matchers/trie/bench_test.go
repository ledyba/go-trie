package trie

import (
	"testing"

	derekparker "github.com/derekparker/trie"
	"github.com/ledyba/go-trie/matchers/test_util"
)

func BenchmarkUnmatchDghubbleTrie(b *testing.B) {
	tr := derekparker.New()
	for _, str := range test_util.Words {
		tr.Add(str, struct{}{})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := range test_util.UnmatchString {
			if tr.HasKeysWithPrefix(test_util.UnmatchString[j:]) {
				b.Error("Should Unmatch!")
			}
		}
	}
}

func BenchmarkUnmatchRegex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if test_util.RegexPattern.MatchString(test_util.UnmatchString) {
			b.Error("Should Unmatch!")
		}
	}
}
func BenchmarkUnmatchTrie(b *testing.B) {
	tr := FromWords(test_util.Words)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if tr.Contains(test_util.UnmatchString) {
			b.Error("Should Unmatch!")
		}
	}
}
