.PHONY: run bench perf build bench-all bench-other bench-go FORCE;

run: bench;

bench: clean
	@make bench-all || make clean

.bin:
	@mkdir -p .bin

.bin/bench: .bin $(shell find . -type f -name *.go)
	@go build -o .bin/bench github.com/ledyba/go-trie/cmds/bench

clean:
	@rm -Rf .bin/ bench

bench-all: bench-other bench-go;

build: .bin/bench;

bench-go: clean .bin/bench FORCE
	@.bin/bench

perf: clean .bin/bench FORCE
	perf stat -e L1-dcache-load-misses -e L1-dcache-loads -e L1-dcache-prefetches .bin/bench
	perf stat -e L1-dcache-load-misses -e L1-dcache-loads -e L1-dcache-prefetches node _rivals/bench.js

bench-other: FORCE
	@php _rivals/bench.php
	@node _rivals/bench.js
