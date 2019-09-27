// Package configuration provides a single place to manage all configuration inputs.
package configuration

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config is the main configuration object.
// It is an immutable, read-only object.
type Config struct {
	vp      *viper.Viper // read from config file
	started time.Time    // date the config object was created
}

// NewConfig constructs a configuration, where
// name is the file name without extension,
// confPath are the path where to look for that file (
// (the "." directory is always included),
// args are the command line arguments (including the command name
// as first argument, as in os.Args).
// If args is nil, then os.Args is used.
// def is a map of default key => values.
// fs is a CFlags structure, containing the accepted flags definitions.
func NewConfig(name string, // file name, no extension
	confPath []string, // file config paths
	args []string, // cli args
	def map[string]interface{}, // default values
	fs *CFlags, // flagset to read from
) *Config {
	c := new(Config)
	c.vp = viper.New()
	c.started = time.Now()

	// parse flags from cd line
	if args == nil || len(args) == 0 {
		fs.Parse(os.Args[1:])
	} else {
		fs.Parse(args[1:])
	}

	// merge flags, only the one visited.
	// If flag is NOT set, the default value does not come from the flag.
	fs.Visit(

		func(f *flag.Flag) {
			k := f.Name
			v, ok := fs.m[k]
			if !ok || v == nil {
				return
			}
			switch t := v.(type) {
			case *int:
				c.vp.Set(k, *t)
			case *float64:
				c.vp.Set(k, *t)
			case *bool:
				c.vp.Set(k, *t)
			case *string:
				c.vp.Set(k, *t)
			default:
				fmt.Printf("\nUnknown type for %v of type %T\n", v, v)
				panic("Type is not implemented yet.")
			}
		})

	// read env data
	c.vp.AutomaticEnv()

	// find and read config file
	c.vp.SetConfigName(name) // Set name without extension
	c.vp.AddConfigPath(".")
	for _, p := range confPath {
		c.vp.AddConfigPath(p)
	}
	err := c.vp.ReadInConfig()
	if err != nil {
		// Continue without a file
		// fmt.Println(err)
	}

	// Set default Viper values for those not set yet
	for k, v := range def {
		c.vp.SetDefault(k, v)
	}

	return c
}

// String human readeable view of the Config.
func (c *Config) String() string {
	var s string
	if f := c.vp.ConfigFileUsed(); len(f) > 0 {
		s += fmt.Sprintf("Using configuration from  %s\n", f)
	} else {
		s += fmt.Sprintf("No configuration file loaded.\n")
	}
	s += fmt.Sprintf("Started : %v ( %v ago)\n",
		c.Started(),
		c.Since())
	s += fmt.Sprintf("--- Env (selected) ---\n")
	for _, k := range []string{"hostname", "home", "pwd", "user"} {
		v := c.vp.GetString(k)
		s += fmt.Sprintf("%s\t= %v\t(%T)\n", k, v, v)
	}

	s += fmt.Sprint("--- Keys & flags ---\n")
	for k, v := range c.vp.AllSettings() {
		s += fmt.Sprintf("%s\t= %v\t(%T)\n", k, v, v)
	}
	return s
}

// Dump the Configuration.
// Used for debugging.
func (c *Config) Dump() *Config {
	fmt.Println(c.String())
	return c
}

// GetInt returns an int.
func (c *Config) GetInt(key string) int {
	return c.vp.GetInt(key)
}

// GetString returns a string
func (c *Config) GetString(key string) string {
	return c.vp.GetString(key)
}

// GetBool returns a bool
func (c *Config) GetBool(key string) bool {
	return c.vp.GetBool(key)
}

// GetFloat64 returns a float64
func (c *Config) GetFloat64(key string) float64 {
	return c.vp.GetFloat64(key)
}

// Get returns any value as an empty interface.
// Typically use for struct or arrays.
func (c *Config) Get(key string) interface{} {
	return c.vp.Get(key)
}

// Started gives time when Config was created.
func (c *Config) Started() time.Time {
	return c.started
}

// Since provides duration since started.
func (c *Config) Since() time.Duration {
	return time.Now().Sub(c.started)
}
