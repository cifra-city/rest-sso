package main

import (
	"os"

	"github.com/cifra-city/rest-sso/internal/cli"
	"github.com/sirupsen/logrus"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
	logrus.Info("Server started successfully")
}
