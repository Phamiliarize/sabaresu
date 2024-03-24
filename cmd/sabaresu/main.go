package main

import (
	"flag"
	"os"

	"github.com/Phamiliarize/sabaresu/pkg/cli"
)

func main() {
	var subcommand string

	runFlags := flag.NewFlagSet("run", flag.ExitOnError)

	cfg := runFlags.String("cfg", "./gateway.toml", "Path to the gateway configuration file")
	schema := runFlags.String("schema", "./schema", "Path to the folder containing schema")
	port := runFlags.Int("port", 3000, "Port to run sabaresu on")

	if len(os.Args) > 1 {
		subcommand = os.Args[1]
	}

	switch subcommand {
	case "init":
		cli.Init()
	case "run":
		runFlags.Parse(os.Args[2:])
		cli.Run(cfg, schema, port)
	default:
		cli.Help()
	}
}
