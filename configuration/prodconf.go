package configuration

import (
	"flag"
	"os"
)

// NewProdConfig defines a production config object.
// Define your own, or adjust this one.
func NewProdConfig() *Config {

	fname := "prodconf"
	fpaths := []string{"./configuration", "../configuration"}

	flags := NewCFlags().
		Add("port", 8080, "port to connect to server").Alias("port", "p").
		Add("host", "localhost", "domain name or ip of server").Alias("host", "h").
		Add("verbose", false, "print verbose information").Alias("verbose", "v").
		Add("debug", false, "activate test and debugging")

	// If go test is running, add some more flags and ignore them
	if flag.Lookup("test.v") != nil {
		flags.
			Add("test.testlogfile", "", "DO NOT USE - Reserved for the go test environment").
			Add("test.v", "", "DO NOT USE - Reserved for the go test environment")
	}

	def := map[string]interface{}{
		"port":    8080,
		"host":    "localhost",
		"debug":   false,
		"verbose": false,
		"version": "0.10",
		"prod":    true,
	}

	return NewConfig(fname, fpaths, os.Args, def, flags)
}
