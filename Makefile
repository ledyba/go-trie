.PHONY: run bench perf build bench-all bench-other bench-go FORCE;

run: bench;

bench: clean
	@make bench-all || make clean

.bin:
	@mkdir -p .bin

.bin/bench: .bin $(shell find . -type f -name *.go)
	@go build -o .bin/bench github.com/ledyba/go-trie/cmds/bench

.bin/bench-regexp: .bin $(shell find . -type f -name *.go)
	@go build -o .bin/bench-regexp github.com/ledyba/go-trie/cmds/bench-regexp

clean:
	@rm -Rf .bin/ bench

bench-all: bench-go bench-other;

build: .bin/bench;

bench-go: clean .bin/bench .bin/bench-regexp FORCE
	@.bin/bench
	@.bin/bench-regexp

perf: clean .bin/bench FORCE
	perf stat -e L1-dcache-load-misses -e L1-dcache-loads -e L1-dcache-prefetches .bin/bench
	perf stat -e L1-dcache-load-misses -e L1-dcache-loads -e L1-dcache-prefetches node _rivals/bench.js

bench-other: FORCE
	@php _rivals/bench.php
	@node _rivals/bench.js
