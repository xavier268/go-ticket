package main

import (
	"github.com/xavier268/go-ticket/app"
	"github.com/xavier268/go-ticket/common/key"
	"github.com/xavier268/go-ticket/configuration"
)

func main() {

	c := configuration.NewProdConfig()

	if c.GetBool(key.VERBOSE) {
		c.Dump()
	}

	a := app.NewApp(c)

	a.Run()

}
