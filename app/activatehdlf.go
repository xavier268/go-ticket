package app

import (
	"net/http"

	"github.com/xavier268/go-ticket/common"
)

// activateHdlr will activate a requested role for the first device that will ask for it.
// Once a request has been received, it cannot be used anymore to activate a role.
// The url should have the request number as its Activate parameter.
func (a *App) activateHdlf(w http.ResponseWriter, r *http.Request) {

	// Check access
	ss := a.Authorize(w, r, common.RoleNone)
	if ss == nil {
		panic("Unexpected nil result for access with RoleNone ?!")
	}

	// Attempt activation
	a.str.Activate(ss.DeviceID, ss.ActReqID)

	// Redirect to absolute ping url  ...
	http.Redirect(w, r, a.CreateAbsoluteURL(a.cnf.API.Ping), http.StatusFound)

}
