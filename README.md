# go-trie

[![Build Status](https://travis-ci.org/ledyba/go-trie.svg?branch=master)](https://travis-ci.org/ledyba/go-trie)
[![Coverage Status](https://coveralls.io/repos/github/ledyba/go-trie/badge.svg?branch=master)](https://coveralls.io/github/ledyba/go-trie?branch=master)

[A trie tree](https://en.wikipedia.org/wiki/Trie) implementation in golang.

It runs faster than [trie to regex](http://google.com/search?q=trie+2+regex) approach more than 100 times (vs golang's regex engine; Please see the benchmark below).

# Performance

Please see [rival](https://github.com/ledyba/go-trie/tree/master/_rivals) directory to benchmarks in other languages.

```bash
 % make bench
 - php's trie2regex: 742.64 us/op (10000 times)
 - v8's trie2regex: 4.29 us/op (100000 times)
 - golang's trie2regex: 2438.09 us / op (1000 times)
 - go-trie: 24.28 us/op (x 100.40) (100000 times)
```

# how to use?

```go
package test

import (
  "testing"
  "github.com/ledyba/go-trie/matchers/trie"
)

func TestSlimReadme(t *testing.T) {
	tr := trie.New() // Animes.
	// kirara
	tr.Add("NewGame!")
	tr.Add("School Live!")
	tr.Add("Urara Meirocho")
	tr.Add("Anne Happy")
	tr.Add("Kiniro Mosaic")
	tr.Add("Hanayamata")
	tr.Add("Is the order a rabbit?")
	tr.Add("Is the order a rabbit??")
	tr.Add("The Demon Girl Next Door")
	tr.Add("Hidamari Sketch")
	tr.Add("Blend S")
	tr.Add("Dōjin Work")
	tr.Add("Magic of Stella")
	// semi-kirara
	tr.Add("Yuki Yuna Is a Hero")
	tr.Add("Non Non Biyori")
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
	if tr.Match("Dojin Work") == true {
		t.Error("Not Dojin Work. \"Dōjin Work\"")
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

You can see other examples in the [unit test](https://github.com/ledyba/go-trie/blob/master/types/trie/trie_test.go) file.