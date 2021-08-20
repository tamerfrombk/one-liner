package main

import (
	"os"
	"github.com/tamerfrombk/one-liner/pkg/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}