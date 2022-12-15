package main

import (
	"bufio"
	"fmt"
	html2text "github.com/pakholeung37/html-to-text"
	"os"
)

func main() {
	fmt.Println("Write down your html:")
	reader := bufio.NewReader(os.Stdin)
	opts := html2text.Options{}
	out, err := html2text.FromReader(reader, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(out)
}
