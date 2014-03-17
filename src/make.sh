#!/bin/bash
set -e

ver="$(cat version)"

tmp="$(readlink -f ../tmp)"
out="$(readlink -f ../out)"

(
	cd oneinit
	go build -o "$tmp/oneinit" -ldflags "-d -X main.version $ver"
)
