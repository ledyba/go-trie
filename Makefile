.PHONY: run bench stat record build bench-all bench-other bench-go FORCE;

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
	@rm -Rf .bin/ bench perf.data perf.data.old

bench-all: bench-go bench-other;

build: .bin/bench;

bench-go: clean .bin/bench .bin/bench-regexp FORCE
	@.bin/bench
	@.bin/bench-regexp

stat: clean .bin/bench FORCE
	perf stat -e branches -e branch-misses -e cache-references -e L1-dcache-load-misses -e L1-dcache-loads -e L1-dcache-prefetches -e cpu-cycles -e instructions -e stalled-cycles-backend -e stalled-cycles-frontend .bin/bench
	perf stat -e branches -e branch-misses -e cache-references -e L1-dcache-load-misses -e L1-dcache-loads -e L1-dcache-prefetches -e cpu-cycles -e instructions -e stalled-cycles-backend -e stalled-cycles-frontend node _rivals/bench.js

record: clean .bin/bench FORCE
	perf record -- .bin/bench
	perf annotate -M intel -M intel-mnemonic

bench-other: FORCE
	@php _rivals/bench.php
	@node _rivals/bench.js
