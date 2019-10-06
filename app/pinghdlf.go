package app

import (
	"net/http"

	"github.com/xavier268/go-ticket/common"
)

// pingHdlf generates a ping page, displaying various informations.
func (a *App) pingHdlf(w http.ResponseWriter, r *http.Request) {

	ss := a.Authorize(w, r, common.RoleNone)
	if ss == nil {
		return
	}

	// send response ... what is shown differes according to Role.
	// See template.
	ss.ExecuteTemplate("ping.html")

}
