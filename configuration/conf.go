// Package configuration provides a single place to manage all configuration inputs.
package configuration

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config is the main configuration object.
// It is an immutable, read-only object.
type Config struct {
	vp *viper.Viper // read from config file
	fs *CFlags      // read from flags
}

// prodDefault is the default production configuration
var prodDefault = map[string]interface{}{
	"port":    8080,
	"host":    "localhost",
	"debug":   false,
	"verbose": false,
}

const prodFileName = "myconf"

// NewProdConfig gets the default production config object.
func NewProdConfig() *Config {

	c := new(Config)
	c.vp = viper.New()
	c.fs = NewTestCFlags()

	// parse flags from cd line
	c.fs.Parse(os.Args[1:])

	// merge flags, only the one visited.
	// If flag is NOT set, the default value does not come from the flag.
	c.fs.Visit(

		func(f *flag.Flag) {
			k := f.Name
			v, ok := c.fs.m[k]
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

	// read config file
	// c.vp.SetConfigFile(prodFileName)
	// c.vp.SetConfigType("yml")
	c.vp.SetConfigName(prodFileName) // Set name without extension
	c.vp.AddConfigPath(".")
	c.vp.AddConfigPath("../configuration/")
	c.vp.AddConfigPath("../cmd/")
	c.vp.AddConfigPath("./cmd/")
	c.vp.AddConfigPath("./go-ticket/cmd/")
	c.vp.AddConfigPath("../go-ticket/cmd/")
	// TODO : add more natural configuration places ...
	err := c.vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	// Set default Viper values for those not set yet
	for k, v := range prodDefault {
		c.vp.SetDefault(k, v)
	}

	return c
}

// String gets a human-readable view of the Config.
func (c *Config) String() string {
	s := fmt.Sprintf("Configuration using file : %s\nHome : %s\nUser : %s\npwd  : %s\nhost0 :%s\n",
		c.vp.ConfigFileUsed(), c.vp.GetString("home"), c.vp.GetString("user"),
		c.vp.GetString("pwd"), c.vp.GetString("host0")) +
		fmt.Sprint("\nKeys   : ")
	for k, v := range c.vp.AllSettings() {
		s += fmt.Sprint(k, "=", v, ", ")
	}
	s += fmt.Sprintf("\nFlags  : ")
	c.fs.Visit(func(f *flag.Flag) {
		s += fmt.Sprint(f.Name, "=", f.Value.String(), ", ")
	})

	return s + "\n"
}

// Dump the Configuration.
// Used for debugging.
func (c *Config) Dump() {
	fmt.Println(c.String())
}

// GetString retrieve a string.
// Priority is flag, then env, then config-file, then default.
func (c *Config) GetString(key string) string {
	r := c.fs.Lookup(key).Value
	if r != nil {
		return r.String()
	}
	return c.vp.GetString(key)
}
