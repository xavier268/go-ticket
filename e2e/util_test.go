package e2e

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xavier268/go-ticket/app"
	"github.com/xavier268/go-ticket/conf"
)

// ========== utilities ===============

// Lauch app asynchroneously - remener to Close it !
func launch() (*app.App, *conf.Conf) {
	c := conf.NewConf()
	a := app.NewApp(c)

	go a.Run()

	if c.Test.Verbose {
		c.Dump()
	}
	return a, c
}

// Get page content as string, or fail on error.
func get(u string) (string, int) {
	r, e := http.Get(u)
	if e != nil {
		panic(e)
	}
	s := r.StatusCode
	if r.StatusCode != http.StatusOK {
		fmt.Println("status : ", r.StatusCode)
		return "", s
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(body), s
}

// get a protected page by providing credentials.
func getWithCredentials(url, user, password string) (string, int) {
	client := new(http.Client)
	rq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	rq.SetBasicAuth(user, password)

	r, e := client.Do(rq)
	if e != nil {
		panic(e)
	}
	s := r.StatusCode
	if r.StatusCode != http.StatusOK {
		fmt.Println("status : ", r.StatusCode)
		return "", s
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(body), s
}
