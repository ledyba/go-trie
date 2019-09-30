.PHONY: bench bench-go bench-php bench-all FORCE;

bench:
	@make bench-all || make clean

.bin:
	@mkdir -p .bin

.bin/bench: .bin
	@go build -o .bin/bench github.com/ledyba/go-trie/cmds/bench

clean:
	@rm -Rf .bin/ bench

bench-go: .bin/bench FORCE
	@.bin/bench

bench-php: FORCE
	@php bench.php

bench-all: bench-php bench-go;
