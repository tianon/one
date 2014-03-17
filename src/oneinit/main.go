package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		print("oneinit: ")
		line, err := r.ReadString('\n')
		if err != nil {
			println(err.Error())
			os.Exit(0)
		} else {
			line = strings.TrimRight(line, "\n")
			println(line)
		}
	}
}
