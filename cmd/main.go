package main

import (
	"fmt"

	"github.com/xavier268/go-ticket/configuration"
)

func main() {
	fmt.Println("Work in progress - testing ...")
	configuration.NewProdConfig().Dump()
	configuration.NewProdConfig().Dump()

}
