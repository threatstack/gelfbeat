package main

import (
	"os"

	"github.com/threatstack/gelfbeat/cmd"

	_ "github.com/threatstack/gelfbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
