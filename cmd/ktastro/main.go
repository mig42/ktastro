package main

import (
	"fmt"
	"os"

	"github.com/mig42/ktastro/internal/cli"
)

func main() {
	err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
