// Copyright 2019-2022 F5 Inc.
// See LICENSE.txt for the Apache-2 License.

package main

// This file is mandatory as otherwise the gelfbeat.test binary is not generated correctly.

import (
	"flag"
	"testing"

	"github.com/threatstack/gelfbeat/cmd"
)

var systemTest *bool

func init() {
	systemTest = flag.Bool("systemTest", false, "Set to true when running system tests")

	cmd.RootCmd.PersistentFlags().AddGoFlag(flag.CommandLine.Lookup("systemTest"))
	cmd.RootCmd.PersistentFlags().AddGoFlag(flag.CommandLine.Lookup("test.coverprofile"))
}

// Test started when the test binary is started. Only calls main.
func TestSystem(t *testing.T) {

	if *systemTest {
		main()
	}
}
