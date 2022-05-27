// Copyright 2019-2022 F5 Inc.
// See LICENSE.txt for the Apache-2 License.

package cmd

import (
	"github.com/threatstack/gelfbeat/beater"

	cmd "github.com/elastic/beats/libbeat/cmd"
	"github.com/elastic/beats/libbeat/cmd/instance"
)

// Name of this beat
var Name = "gelfbeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: Name})
