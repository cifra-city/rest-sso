package main

import (
	"os"

	"github.com/cifra-city/rest-sso/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
