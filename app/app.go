// Package app provides server services common to both admin and exploitation.
package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xavier268/go-ticket/common"
	"github.com/xavier268/go-ticket/configuration"
	"github.com/xavier268/go-ticket/configuration/key"
)

// App is the application server.
type App struct {
	srv *http.Server          // server
	cnf *configuration.Config // config
	str common.Store          // data store
}

// NewApp constructs  a new AppServer.
// It is configured from the provided configuration.
func NewApp(c *configuration.Config) *App {

	a := &App{new(http.Server), c, nil}
	a.srv.Addr = c.GetString(key.ADDR)

	// dump config if verbose
	if c.GetBool(key.VERBOSE) {
		c.Dump()
	}

	// Set  handlers in a new mux
	mux := http.NewServeMux()
	mux.Handle("/ping/", a.messageHdlr("pong !"))

	// Save mux in server
	a.srv.Handler = mux

	return a
}

// Run the app
func (a *App) Run() error {
	if a.cnf.GetBool(key.VERBOSE) {
		fmt.Println("Listening on ", a.cnf.GetString(key.ADDR))
	}
	return a.srv.ListenAndServe()
}

// messageHdlr returns  a message sending handler.
// A set/read cookie
func (a *App) messageHdlr(txt string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Message to be sent : ", txt)
			// read cookie
			fmt.Println("Read cookie : ", readCookie(r, "moncook"))
			// set cookie
			setCookie(w, "moncook", "ma valeur", 100*time.Second)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, txt)
		},
	)
}

// tracingMdlw is a tracing middleware for debugging.
func (a *App) tracingMdlw(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("About to serve request ...")
			next.ServeHTTP(w, r)
			fmt.Println(".... request was served")
		},
	)
}

// Close the App.
func (a *App) Close() error {
	if a.srv != nil {
		a.srv.Close()
	}
	if a.str != nil {
		a.str.Close()
	}
	return nil
}
