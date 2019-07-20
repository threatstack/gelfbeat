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

