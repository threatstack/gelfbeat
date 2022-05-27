// gelfbeat
// Copyright 2019-2022 F5 Inc.
// See LICENSE.txt for the Apache-2 License.

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
