package app

import (
	"fmt"
	"net/http"
)

// activateHdlr will activate a requested role for the first device that will ask for it.
// Once a request has been received, it cannot be used anaymore to activate a role.
// The url should have the request number as its 'a' query parameter.
func (a *App) activateHdlf(w http.ResponseWriter, r *http.Request) {

	did := a.getDeviceID(w, r)
	reqid := r.URL.Query().Get("a")
	role, err := a.str.Activate(did, reqid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><h1>Activation success<h1>Role is now : %s</html>",
		role.String())
}
