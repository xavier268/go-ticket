package app

import (
	"net/http"

	"github.com/xavier268/go-ticket/common"
)

// HandlerFuncuntion to logout, back to RoleNone
func (a *App) logoutHdlf(w http.ResponseWriter, r *http.Request) {

	ss := a.Authorize(w, r, common.RoleNone)
	a.str.UnsetRole(ss.DeviceID)

	// Redirect to ping ...
	http.Redirect(w, r, a.CreateAbsoluteURL(a.cnf.API.Ping), http.StatusFound)

}
