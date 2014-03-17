package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var version = "UNK"

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("oneinit v%s: ", version)
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(0)
		} else {
			line = strings.TrimRight(line, "\n")
			fmt.Println(line)
		}
	}
}
