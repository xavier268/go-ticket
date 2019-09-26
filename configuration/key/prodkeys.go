// Package key contains key constants for flags and configuration.
// Use constant instead of string to prevent misspelling !
package key

// Configuration keys used from the configuration object.
// Note that all keys are lower_case.
const (
	HOME    = "home"
	USER    = "user"
	PORT    = "port"
	ADDR    = "addr"
	TIMEOUT = "timeout" // time out in seconds
	VERSION = "version"
	PWD     = "pwd" // current working dir
	VERBOSE = "verbose"
)

// TESTFLAGS are the flags that are added during go test and that should be ignored.
var TESTFLAGS = []string{"test.testlogfile", "test.v"}
