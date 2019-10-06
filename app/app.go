// Package app provides server services common to both admin and exploitation.
package app

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/xavier268/go-ticket/common"
	"github.com/xavier268/go-ticket/conf"
	"github.com/xavier268/go-ticket/impl/barcode"
	"github.com/xavier268/go-ticket/impl/memstore"
)

// App is the application server.
type App struct {
	srv  *http.Server    // server
	cnf  *conf.Conf      // config
	str  common.Storer   // data store
	rand *rand.Rand      // random generator
	bc   common.BarCoder // barcode encoding
}

// NewApp constructs  a new AppServer.
// It is configured from the provided configuration.
func NewApp(c *conf.Conf) *App {

	a := new(App)
	a.srv = new(http.Server)
	a.srv.Addr = c.Addr.Private
	a.cnf = c
	a.str = memstore.New()
	a.rand = rand.New(rand.NewSource(time.Now().UnixNano() + 111111111)) // initialize random gen
	a.bc = barcode.New()
	a.bc.SetFormat(a.cnf.Barcode.Format)
	// dump config if verbose
	if c.Test.Verbose {
		c.Dump()
	}

	// Set  handlers in a new mux
	mux := http.NewServeMux()
	mux.HandleFunc(c.API.QRImage, a.qrHdlf)
	mux.HandleFunc(c.API.Activate, a.activateHdlf)
	mux.HandleFunc(c.API.Admin, a.adminHdlf)
	mux.HandleFunc(c.API.Ping, a.pingHdlf)
	mux.HandleFunc(c.API.Ticket, a.tktHdlf)
	mux.HandleFunc(c.API.Logout, a.logoutHdlf)
	mux.HandleFunc("/", a.h404)

	// Save mux in server
	a.srv.Handler = mux

	// Set root password if not set
	if len(a.cnf.Superuser.Password) < 8 {
		a.cnf.Superuser.Password = strconv.FormatInt(a.rand.Int63(), 36)
		fmt.Printf("Changed root password to : %s\n", a.cnf.Superuser.Password)
	}

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

// h404 displays 404 not found page.
func (a *App) h404(w http.ResponseWriter, r *http.Request) {
	ss := a.Authorize(w, r, common.RoleNone)
	ss.ExecuteTemplate("404.html")
}
