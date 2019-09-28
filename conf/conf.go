// Package conf concentrate all the configuration plumbing.
package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

// Conf defines the structure of a configuration object.
type Conf struct {
	Start  time.Time // Date started
	Parsed struct {  // What exactly did we read yet ?
		Default bool
		File    bool
		Env     bool
		Flags   bool
	}
	Version string // Version

	Command string   // Command name typed
	Args    []string // Non-flag arguments

	File struct {
		Name  string   // Conf file, with json extension
		Used  string   // File actually used, full name
		Paths []string // Path to look for conf file
	}

	Env struct {
		Home string // env
		User string // env
		PWD  string // env
	}

	Test struct {
		Verbose bool   // Verbose
		Short   bool   // Short tests
		LogFile string // not used. From go test.
	}

	Addr struct {
		Public  string // External addr (eg http://jhg.com:80)
		Private string // Internal addr (eg :8080)
	}
	Cookie struct {
		Name   string // name of devide id cookie
		MaxAge int    // in seconds
	}
}

// NewConf constructs and parse a new Conf object.
// For testing purposes, you may alter it afterwards.
func NewConf() *Conf {
	c := new(Conf)
	c.loadDefault()
	c.loadFile()
	c.loadEnv()
	c.loadFlags(os.Args)
	return c
}

// Dump prints a human readable config.
func (c *Conf) Dump() *Conf {
	fmt.Printf("\n%v\n", *c)
	return c
}

// loadDefault set the defaults.
func (c *Conf) loadDefault() {
	c.Version = "0.14"
	c.Start = time.Now()

	c.Parsed.Default = true

	c.File.Name = "go-ticket.json"
	c.File.Paths = append(c.File.Paths, ".", "../conf", "./conf")

	c.Addr.Private = ":8080"
	c.Cookie.Name = "deviceid"
	c.Cookie.MaxAge = 3600 * 24 // 24 h
}

// loadFile read the config file, json format only.
func (c *Conf) loadFile() {
	if len(c.File.Name) == 0 {
		panic("Conf File name was not set ?!")
	}
	if path.Ext(c.File.Name) != ".json" {
		panic("Conf File Name should have the .json extension.")
	}
	fmt.Printf("Looking for '%s' in %#v\n", c.File.Name, c.File.Paths)
	for _, p := range c.File.Paths {
		fn := path.Join(p, c.File.Name)
		fmt.Println("Trying ", fn)
		content, err := ioutil.ReadFile(fn)
		if err == nil {
			// Found ! process ...
			// c.Dump()
			fmt.Println("Loading ", fn)
			err = json.Unmarshal(content, c)
			// c.Dump()
			if err != nil {
				fmt.Println("*** Error reading config file ", fn, "\n*** ", err)
				return
			}
			// Do not search for other files.
			c.File.Used = fn
			c.Parsed.File = true
			return
		}
	}
	fmt.Println("No configuration file found")
}

// loadEnv reads selected environement values.
func (c *Conf) loadEnv() {

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		k := strings.ToUpper(pair[0])
		v := strings.Join(pair[1:], "=")
		switch k {
		case "HOME":
			c.Env.Home = v
		case "USER":
			c.Env.User = v
		case "PWD":
			c.Env.PWD = v
		}
	}
	c.Parsed.Env = true

}

// loadFlags parse flags from the provided []string.
// The first element in the array is the command name.
func (c *Conf) loadFlags(from []string) {

	if len(from) == 0 {
		// Nothing to parse.
		return
	}
	c.Command = from[0]
	fs := flag.NewFlagSet(c.Command, flag.ExitOnError)

	// Define relevant flags
	fs.BoolVar(&c.Test.Verbose, "verbose", c.Test.Verbose, "Verbose execution.")
	fs.BoolVar(&c.Test.Verbose, "v", c.Test.Verbose, "Verbose execution.")
	fs.BoolVar(&c.Test.Verbose, "test.v", c.Test.Verbose, "Verbose execution.")
	fs.BoolVar(&c.Test.Verbose, "test.verbose", c.Test.Verbose, "Verbose execution.")

	fs.BoolVar(&c.Test.Short, "test.short", c.Test.Short, "Verbose execution.")
	fs.BoolVar(&c.Test.Short, "short", c.Test.Short, "Verbose execution.")

	fs.StringVar(&c.Test.LogFile, "test.testlogfile", c.Test.LogFile, "DO NOT USE - used by go test.")

	fs.StringVar(&c.Addr.Public, "pubaddr", c.Addr.Public, "Public address.")
	fs.StringVar(&c.Addr.Private, "privaddr", c.Addr.Private, "Private address.")

	// Actual parsing
	fs.Parse(from[1:])

	// Remaining args ..
	c.Args = fs.Args()

	// Display default flags.
	// fs.PrintDefaults()

	fmt.Println("Flag parsed : ", *fs)

	c.Parsed.Flags = true
}
