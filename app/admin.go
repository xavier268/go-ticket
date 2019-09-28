package app

import (
	"fmt"
	"net/http"
)

// Display admin page.
// Require at least Admin role to get in.
func (a *App) adminHdlf(w http.ResponseWriter, r *http.Request) {

	did := a.getDeviceID(w, r)
	role := a.str.GetRole(did)

	/* if role < common.RoleAdmin {
		w.WriteHeader(http.StatusForbidden)
		return
	} */
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<http><h1>Admin</h1>Role : %s</http>", role.String())

}
