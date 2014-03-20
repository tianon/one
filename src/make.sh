#!/bin/bash
set -e

ver="$(cat version)"

tmp="$(readlink -f ../tmp)"
out="$(readlink -f ../out)"

(
	set -x
	cd oneinit
	go build -o "$tmp/oneinit" -ldflags "-d -X main.version $ver"
)

(
	set -x
	cd ../../linux
	make
)

cp -v ../../linux/arch/x86/boot/bzImage "$tmp/linux"

# TODO create initrd.gz
# TODO create hybrid ISO
