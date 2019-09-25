package configuration

import (
	"testing"
)

func TestConfigFromFile(t *testing.T) {

	c := NewConfig("config_test", // file name
		nil,                     // file paths
		[]string{"testCommand"}, // cli args
		nil,                     // default keys
		NewCFlags(),             // Empty CFlags
	)
	//c.Dump()
	if len(c.GetString("home")) == 0 ||
		len(c.GetString("pwd")) == 0 ||
		len(c.GetString("user")) == 0 ||

		c.GetFloat64("srv.test") != 3.14 ||
		c.GetInt("srv.deux") != 2 ||

		c.GetInt("port") != 1234 ||
		c.GetInt("PORT") != 1234 ||
		c.GetInt("PoRt") != 1234 ||

		c.GetBool("serv.nope") != false ||
		c.GetString("nope") != "" ||
		c.GetInt("nope") != 0 ||
		c.GetFloat64("nope") != 0. ||

		false {
		c.Dump()
		t.FailNow()
	}
	// testing Get for a complex struct,
	// it takes some jungling with type conversion, though ...
	srv := c.Get("srv").(map[string]interface{})
	arr := srv["arr"].([]interface{})
	if arr[1].(int) != 21 {
		c.Dump()
		t.Fatal(srv, arr, " - Failed access inside a complex struct read from file")
	}
}

func TestConfigWithDefaults(t *testing.T) {

	c := NewConfig("config_test", // file name
		nil,                     // file paths
		[]string{"testCommand"}, // cli args
		map[string]interface{}{ // default keys
			"srv.def":   8.88,
			"port":      8888,
			"defstring": "this is a default string",
			"defint":    888,
			"deffloat":  8.88,
		},
		NewCFlags(),
	)
	//c.Dump()
	if len(c.GetString("home")) == 0 ||
		len(c.GetString("pwd")) == 0 ||
		len(c.GetString("user")) == 0 ||

		c.GetFloat64("srv.test") != 3.14 ||
		c.GetInt("srv.deux") != 2 ||

		c.GetInt("port") != 1234 ||
		c.GetInt("PORT") != 1234 ||
		c.GetInt("PoRt") != 1234 ||

		c.GetBool("srv.nope") != false ||
		c.GetString("nope") != "" ||
		c.GetInt("nope") != 0 ||

		c.GetString("defstring") != "this is a default string" ||
		c.GetFloat64("deffloat") != 8.88 ||
		c.GetInt("defint") != 888 ||
		c.GetFloat64("srv.def") != 8.88 ||

		false {
		c.Dump()
		t.FailNow()
	}
}
func TestConfigWithDefaultsAndFlags(t *testing.T) {

	fgs := NewCFlags().
		Add("port", 9999, "port").Alias("port", "p").
		Add("host", "default host from flag", "host name")

	c := NewConfig("config_test", // file name
		nil, // file paths
		[]string{"testCommand", "-port=777", "-host", "host from cli", "non flags"}, // cli args
		map[string]interface{}{ // default keys
			"srv.def":   8.88,
			"port":      8888,
			"defstring": "this is a default string",
			"defint":    888,
			"deffloat":  8.88,
		},
		fgs,
	)
	// c.Dump()
	if len(c.GetString("home")) == 0 ||
		len(c.GetString("pwd")) == 0 ||
		len(c.GetString("user")) == 0 ||

		c.GetFloat64("srv.test") != 3.14 ||
		c.GetInt("srv.deux") != 2 ||

		c.GetInt("port") != 777 ||
		c.GetInt("PORT") != 777 ||
		c.GetInt("PoRt") != 777 ||

		c.GetBool("srv.nope") != false ||
		c.GetString("nope") != "" ||
		c.GetInt("nope") != 0 ||

		c.GetString("defstring") != "this is a default string" ||
		c.GetFloat64("deffloat") != 8.88 ||
		c.GetInt("defint") != 888 ||
		c.GetFloat64("srv.def") != 8.88 ||

		c.GetString("host") != "host from cli" ||

		false {
		c.Dump()
		t.FailNow()
	}
}

func TestProdConfig(t *testing.T) {
	c := NewProdConfig() // .Dump()
	if !c.GetBool("prod") {
		c.Dump()
		t.Fatal("The prod key is not set ?!")
	}
	if len(c.vp.ConfigFileUsed()) == 0 {
		c.Dump()
		t.Fatal("No prod configuration file found !?")
	}

}
