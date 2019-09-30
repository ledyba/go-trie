# go-trie

A trie tree implementation in git.

It runs faster than [trie to regex]() approach more than 50 times (Please see the benchmark below).

## how to use?

[Please see test codes]().

# Result

```bash
 % make bench
[Baseline] benchmarking php's regexp (10000 times).
 - result: 568.002391 us / op

[Baseline] benchmarking golang's regexp (1000 times).
 - result: 2379.605000 us / op

[Ours] benchmarking trie (100000 times).
 - result: 37.644730 us / op
2019/09/30 16:09:05 Trie is faster 63.21 fastar than regex!
```