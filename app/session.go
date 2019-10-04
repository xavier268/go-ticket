package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/xavier268/go-ticket/common"
)

// SessionData is the (maximum) session information structure.
type SessionData struct {
	DeviceID                  string      // from cookie
	Role                      common.Role // from database
	TicketID, ActReqID, QRTxt string      // from Query Params
	CredentialsUser           string      // From basic auth
	CredentialsProvided       bool        // From basic auth
	CredentialsValid          bool        // From basic auth
}

// Authorize provides access to session parameters and manages authorization.
// MinimalRole defines the minimal expected role level.
// If not authorized, sent Unauthorized header and return nil.
func (a *App) Authorize(w http.ResponseWriter, r *http.Request, minimalRole common.Role) *SessionData {

	ss := new(SessionData)

	// first, get/set session cookie.

	name := a.cnf.Cookie.Name
	age := a.cnf.Cookie.MaxAge
	c, err := r.Cookie(name)
	if err == nil {
		// cookie  is set .. use its value
		ss.DeviceID = c.Value
	} else {
		// cookie not set, ... create a value
		ss.DeviceID = strconv.FormatInt(a.rand.Int63(), 36)
	}
	// Don't attempt to reuse read cookie,
	// that creates strange issues with the Path ...
	c = new(http.Cookie)
	c.Name = name
	c.Value = ss.DeviceID
	// In all cases, rewrite cookie, updating age, expires and path.
	c.Path = "/"
	c.MaxAge = age
	c.Expires = time.Now().Add(time.Duration(age) * time.Second)
	// fmt.Println("\tCookie sent : ", c)
	http.SetCookie(w, c)

	// look for other information : role and various query params

	ss.Role = a.str.GetRole(ss.DeviceID)

	ss.TicketID = r.URL.Query().Get(a.cnf.API.QueryParam.Ticket)
	ss.ActReqID = r.URL.Query().Get(a.cnf.API.QueryParam.ActivationRequestID)
	ss.QRTxt = r.URL.Query().Get(a.cnf.API.QueryParam.QRText)

	// Read credentials if provided

	p := ""
	ss.CredentialsUser, p, ss.CredentialsProvided = r.BasicAuth()
	ss.CredentialsValid = ss.CredentialsProvided &&
		ss.CredentialsUser == a.cnf.Superuser.Name &&
		p == a.cnf.Superuser.Password

	if ss.CredentialsProvided && ss.CredentialsValid {
		ss.Role = common.RoleSuper
	}

	// Current authorization is below expectation,
	// Ask (again) for authentication
	if ss.Role < minimalRole {
		w.Header().Add("WWW-Authenticate", "Basic realm=go-ticket")
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	return ss

}
