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

	u, p, ok := r.BasicAuth()
	fmt.Printf("Provided credentials : %s  %s %T\n", u, p, ok)
	fmt.Println()

	// Check credentials if provided.
	ok = ok && u == a.cnf.Superuser.Name && p == a.cnf.Superuser.Password

	if !ok {
		// Ask (again) for authentication
		w.Header().Add("WWW-Authenticate", "Basic realm=go-ticket")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><h1>Admin</h1>Role : %s</html>", role.String())

}
