# go-trie

A trie tree implementation in git.

## how to use?

[Please see test codes]().

# Result

```bash
 % make bench
[Baseline] benchmarking php's regexp (10000 times).
 - result: 576.267004 us / op

[Baseline] benchmarking golang's regexp (1000 times).
 - result: 2348.964000 us / op

[Ours] benchmarking trie (100000 times).
 - result: 36.620080 us / op

2019/09/30 16:00:40 Trie is faster 64.14 fastar than regex!
```