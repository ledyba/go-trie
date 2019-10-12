# go-trie

[![Build Status](https://travis-ci.org/ledyba/go-trie.svg?branch=master)](https://travis-ci.org/ledyba/go-trie)
[![Coverage Status](https://coveralls.io/repos/github/ledyba/go-trie/badge.svg?branch=master)](https://coveralls.io/github/ledyba/go-trie?branch=master)

[A trie tree](https://en.wikipedia.org/wiki/Trie) implementation in golang.

It runs faster than [trie to regex](http://google.com/search?q=trie+2+regex) approach more than 100 times (vs golang's regex engine; Please see the benchmark below).

# Performance

Please see [rival](https://github.com/ledyba/go-trie/tree/master/_rivals) directory to benchmarks in other languages.

## benchmark

```bash
 % make bench
 -             go-trie: 18.60   us/op (100000 times)
 - golang's trie2regex: 2336.16 us/op (  1000 times)
 -    php's trie2regex: 849.18  us/op ( 10000 times)
 -     v8's trie2regex: 4.25    us/op (100000 times)
```

## zero heap allocation

`Contains(string)` and `Match(string)` operation do not allocate heaps:

```bash
% 
goos: linux
goarch: amd64
pkg: github.com/ledyba/go-trie/matchers/trie
BenchmarkUnmatchTrie-32                64381       18400 ns/op         0 B/op        0 allocs/op
PASS
ok    github.com/ledyba/go-trie/matchers/trie 4.086s

```

# how to use?

```go
package test

import (
  "testing"
  "github.com/ledyba/go-trie/matchers/trie"
)

func TestReadme(t *testing.T) {
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

You can see other examples in the [unit test](https://github.com/ledyba/go-trie/blob/master/matchers/trie/trie_test.go) file.

# LICENSE

Licensed under 2-clause BSD license.

This software uses materials from the wikipedia article [五稜郭](https://ja.wikipedia.org/wiki/%E4%BA%94%E7%A8%9C%E9%83%AD) and [ポケモン一覧](https://ja.wikipedia.org/wiki/%E3%83%9D%E3%82%B1%E3%83%A2%E3%83%B3%E4%B8%80%E8%A6%A7), which is released under the <a href="https://creativecommons.org/licenses/by-sa/3.0/">Creative Commons Attribution-Share-Alike License 3.0</a>.
