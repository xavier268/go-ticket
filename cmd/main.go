package main

import (
	"github.com/xavier268/go-ticket/app"
	"github.com/xavier268/go-ticket/conf"
)

func main() {

	c := conf.NewConf()

	if c.Test.Verbose {
		c.Dump()
	}

	a := app.NewApp(c)

	a.Run()

}
