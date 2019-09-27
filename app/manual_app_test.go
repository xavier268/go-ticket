package app

import (
	"testing"
	"time"

	"github.com/xavier268/go-ticket/configuration"
	"github.com/xavier268/go-ticket/configuration/key"
)

func getTestConfig() *configuration.Config {
	return configuration.NewConfig(
		"testapp",
		[]string{},
		[]string{"testing"},
		map[string]interface{}{
			key.ADDR:       ":8080",
			key.PUBLICADDR: "http://192.168.1.9",
			key.VERBOSE:    true},
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
	//t.Skip()
	a := NewApp(getTestConfig())
	go a.Run() // run server for 10 seconds, then close it ...
	time.Sleep(2000 * time.Second)
	a.Close()
}
