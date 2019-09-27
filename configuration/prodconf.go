package configuration

import (
	"os"

	"github.com/xavier268/go-ticket/common/key"
)

// NewProdConfig defines a production config object.
// Define your own, or adjust this one.
func NewProdConfig() *Config {

	fname := "go-ticket"
	fpaths := []string{"./cmd", "../cmd",
		"./configuration", "../configuration",
		"../../etc", "../etc", "/etc/"}

	flags := NewCFlags()

	// Add test flags, in case we test ;-)
	for _, k := range key.TESTFLAGS {
		flags.Add(k, "", "DO NOT USE - Reserved for the go test configuration.")
	}

	// Then, add actual production flags.
	// NO ALIAS in production (buggy, needs fixing)
	flags.
		Add(key.ADDR, "", "full address URL to run the server").
		Add(key.VERBOSE, false, "print verbose information")

	// Define defaults key values.
	def := map[string]interface{}{
		key.VERBOSE:   false,
		key.VERSION:   "0.12",
		key.PROD:      true,
		key.COOKIEKEY: "did",
		key.COOKIEAGE: 3600 * 24 * 15, // 15 jours
	}

	return NewConfig(fname, fpaths, os.Args, def, flags)
}
