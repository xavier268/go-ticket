package app

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/xavier268/go-ticket/common/key"
)

// getDeviceID gets the device ID from the request.
// If no device ID set yet, create one and set it as a cookie in the response headers.
func (a *App) getDeviceID(w http.ResponseWriter, r *http.Request) string {
	name := a.cnf.GetString(key.COOKIEKEY)
	age := a.cnf.GetInt(key.COOKIEAGE)
	var v string
	verbose := a.cnf.GetBool(key.VERBOSE)

	c, err := r.Cookie(name)

	if verbose {
		fmt.Println("Cookie read : ", c)
	}

	if err == nil {
		// cookie  set .. use value
		v = c.Value
	} else {
		// cookie not set, ... create value
		v = strconv.FormatInt(a.rand.Int63(), 36)
	}
	// Don't attempt to reuse read cookie,
	// that creates strange issues with the Path ...
	c = new(http.Cookie)
	c.Name = name
	c.Value = v
	// In all cases, rewrite cookie, updating age, expires and path.
	c.Path = "/"
	c.MaxAge = age
	c.Expires = time.Now().Add(time.Duration(age) * time.Second)
	// fmt.Println("\tCookie sent : ", c)
	http.SetCookie(w, c)
	return c.Value
}
