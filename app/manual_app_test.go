package app

import (
	"testing"
	"time"

	"github.com/xavier268/go-ticket/conf"
)

// gettestConfig provides a test configuration.
func getTestConfig() *conf.Conf {
	return conf.NewConf()
}

func TestDump(t *testing.T) {
	getTestConfig().Dump()
}

func TestApp(t *testing.T) {
	a := NewApp(getTestConfig())
	go a.Run() // run server for 5 minutes, then close it ...
	if testing.Short() {
		time.Sleep(5 * time.Second)
	} else {
		time.Sleep(15 * time.Minute)
	}
	a.Close()
}
