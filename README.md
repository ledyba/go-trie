# go-trie

[A trie tree](https://en.wikipedia.org/wiki/Trie) implementation in golang.

It runs faster than [trie to regex]() approach more than 50 times (Please see the benchmark below).

# Performance

please see [rival]() directory to benchmarks in other languages.

```bash
 % make bench
 - php's regexp (10000 times): 583.44 us / op
 - v8's regexp (100000 times): 4.22 us / op
 - golang's regexp (1000 times): 2415.712000 us / op
 - go-trie (10000 times): 38.492500 us / op (x 62.76)

```

# how to use?

```go
package test

import (
  trie "github.com/ledyba/go-trie"
)

func TestReadme(t *testing.T) {
  tr := trie.New() // Animes.
  tr.Add("NewGame!")
  tr.Add("School Live!")
  tr.Add("Urara Meiro Chou")
  tr.Add("Yuki Yuna Is a Hero")
  tr.Add("Non Non Biyori.")
  tr.Add("Anne Happy")
  tr.Add("Kiniro Mosaic")
  tr.Pack()

  // Match method
  if tr.Match("NewGame!") == false {
    t.Error("NewGame! is a first season of the series.")
  }
  if tr.Match("NewGame!!") == false {
    t.Error("NewGame!! is a second season of the series.")
  }
  if tr.Match("NewGame") == true {
    t.Error("Not NewGame. NewGame\"!\"")
  }

  // Contains method
  if tr.Contains("I would like to eat udon with Fuu Inubozaki, a hero in \"Yuki Yuna Is a Hero\".") == false {
    t.Error("What????? Why????")
  }
  if tr.Contains("Alas, Ikaruga is going...") == true {
    t.Error("Ikaruga is a game. Not an animation.")
  }
}
```

You can see other examples in [unit test]() file.