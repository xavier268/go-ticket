// Package app provides server services common to both admin and exploitation.
package app

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/xavier268/go-ticket/common"
	"github.com/xavier268/go-ticket/conf"
)

// App is the application server.
type App struct {
	srv  *http.Server // server
	cnf  *conf.Conf   // config
	str  common.Store // data store
	rand *rand.Rand   // random generator
}

// NewApp constructs  a new AppServer.
// It is configured from the provided configuration.
func NewApp(c *conf.Conf) *App {

	a := new(App)
	a.srv = new(http.Server)
	a.srv.Addr = c.Addr.Private
	a.cnf = c
	a.str = nil                                              // TODO
	a.rand = rand.New(rand.NewSource(time.Now().UnixNano())) // initialize random gen

	// dump config if verbose
	if c.Test.Verbose {
		c.Dump()
	}

	// Set  handlers in a new mux
	mux := http.NewServeMux()
	mux.HandleFunc("/ping/", a.pingHdlf)
	mux.HandleFunc("/qr/", a.qrHdlf)

	// Save mux in server
	a.srv.Handler = mux

	return a
}

// Run the app
func (a *App) Run() error {
	if a.cnf.Test.Verbose {
		fmt.Println("Listening privately on ", a.cnf.Addr.Private)
		fmt.Println("Listening publicly on ", a.cnf.Addr.Public)
	}
	return a.srv.ListenAndServe()
}

// Close the App.
// Server is closed using shutdown, with a 1 mn timeout.
func (a *App) Close() error {
	if a.srv != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()
		a.srv.Shutdown(ctx)
	}
	if a.str != nil {
		a.str.Close()
	}
	return nil
}
