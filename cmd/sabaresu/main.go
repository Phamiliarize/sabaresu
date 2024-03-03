package main

import (
	"os"

	"github.com/Phamiliarize/sabaresu/pkg/cli"
)

func main() {
	// Simple CLI interface
	switch os.Args[1] {
	case "init":
		cli.Init()
	case "run":
		cli.Run()
	default:
		cli.Help()
	}
}
