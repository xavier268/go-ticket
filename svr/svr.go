// Package svr provides server services common to both admin and exploitation.
package svr

import (
	"net/http"

	"github.com/xavier268/go-ticket/configuration"
)

// AppServer is the application server.
type AppServer struct {
	*http.Server
	c *configuration.Config
}

// NewAppServer constructs  a new AppServer.
// It is configured from the provided configuration.
func NewAppServer(c *configuration.Config) *AppServer {
	a := &AppServer{new(http.Server), c}
	return a
}
