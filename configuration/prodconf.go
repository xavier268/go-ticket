package configuration

import (
	"os"

	"github.com/xavier268/go-ticket/configuration/key"
)

// NewProdConfig defines a production config object.
// Define your own, or adjust this one.
func NewProdConfig() *Config {

	fname := "prodconf"
	fpaths := []string{"./configuration", "../configuration"}

	flags := NewCFlags()

	// Add test flags, in case we test ;-)
	for _, k := range key.TESTFLAGS {
		flags.Add(k, "", "DO NOT USE - Reserved for the go test configuration.")
	}

	// Then, add actual production flags.
	flags.
		Add(key.ADDR, "localhost", "domain name or ip of server").Alias(key.ADDR, "a").
		Add(key.VERBOSE, false, "print verbose information").Alias(key.VERBOSE, "v")

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
