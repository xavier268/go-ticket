// Package key contains key constants for flags and configuration.
// Use constant instead of string to prevent misspelling !
package key

// Configuration keys used from the configuration object.
// Note that all keys are lower_case.
const ( //      keys ...       // are pointing to ...
	HOME       = "home"        // string
	USER       = "user"        // string
	ADDR       = "addr"        // string, address includes port, such as "http:lklkj.com:8080"
	PUBLICADDR = "public.addr" // string, addr as needed by the public from outside
	TIMEOUT    = "timeout"     // int, time out in seconds
	VERSION    = "version"     // string
	PWD        = "pwd"         // string, current working dir
	VERBOSE    = "verbose"     // bool
	COOKIEKEY  = "cookiekey"   // string, Key for the device id cookie
	COOKIEAGE  = "cookieage"   // int, Default max-age in seconds for the device id cookie.
)

// TESTFLAGS are the flags that are added during go test and that should be ignored.
var TESTFLAGS = []string{"test.testlogfile", "test.v"}
