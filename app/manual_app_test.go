package app

import (
	"testing"
	"time"

	"github.com/xavier268/go-ticket/configuration"
	"github.com/xavier268/go-ticket/configuration/key"
)

// gettestConfig provides a test configuration.
func getTestConfig() *configuration.Config {
	return configuration.NewConfig(
		"testapp",
		[]string{},
		[]string{"testing"},
		map[string]interface{}{
			key.ADDR:       ":8080",
			key.PUBLICADDR: "http://192.168.1.9",
			key.VERBOSE:    true,
			key.COOKIEKEY:  "did",
			key.COOKIEAGE:  3600 * 24, //  seconds default for device id before renewal
		},
		configuration.NewCFlags(),
	)
}

func TestConfigOK(t *testing.T) {
	c := getTestConfig()
	if c.GetString(key.ADDR) != ":8080" || !c.GetBool(key.VERBOSE) {
		t.Fatal("Test configuration did not work ?")
	}
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
