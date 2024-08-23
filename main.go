package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	var nFlag = flag.Int("n", 1, "number of times to print the message")
	flag.Parse()
	fmt.Printf("Hello, World! %d\n", *nFlag)
	_, _ = git.PlainClone("./xxx", false, &git.CloneOptions{
		URL: "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
}
