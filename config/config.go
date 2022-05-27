// Copyright 2019-2022 F5 Inc.
// See LICENSE.txt for the Apache-2 License.

// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Listen string `config:"listen"`
	Port   int    `config:"port"`
}

var DefaultConfig = Config{
	Listen: "127.0.0.1",
	Port:   12201,
}
