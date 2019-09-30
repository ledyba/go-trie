.PHONY: run bench bench-all FORCE;

run: bench;

bench: clean
	@make bench-all || make clean

.bin:
	@mkdir -p .bin

.bin/bench: .bin
	@go build -o .bin/bench github.com/ledyba/go-trie/cmds/bench

clean:
	@rm -Rf .bin/ bench

bench-all: .bin/bench FORCE
	@php others/bench.php
	@echo ""
	@node others/bench.js
	@echo ""
	@.bin/bench
