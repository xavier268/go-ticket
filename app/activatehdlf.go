package app

import (
	"fmt"
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

	// Attempt activate
	role, err := a.str.Activate(ss.DeviceID, ss.ActReqID)

	if err != nil {
		fmt.Fprintf(w, "<html><h1>Activation failure</h1> Role is still : %s</html>",
			role.String())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><h1>Activation success</h1>Role is now : %s</html>",
		role.String())
}
