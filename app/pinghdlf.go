package app

import (
	"fmt"
	"net/http"
)

// pingHdlf generates a ping page, displaying various informations.
func (a *App) pingHdlf(w http.ResponseWriter, r *http.Request) {

	// Get/Set device id and corresponding role
	did := a.getDeviceID(w, r)
	role := a.str.GetRole(did)

	// send response ...
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "<html><h1>Ping response</h1>")

	fmt.Fprintf(w, "\n<br/><h2>Request</h2> <br/>Url : %s<br/>Device id : %s<br/>Role : %s",
		r.URL, did, role.String())

	fmt.Fprintf(w, "\n<h2>Headers</h2>")
	for k, v := range r.Header {
		fmt.Fprintf(w, "\n<br/>%v : %v", k, v)
	}

	fmt.Fprintf(w, "\n<br/><h2>Cookies</h2>")
	for _, c := range r.Cookies() {
		fmt.Fprintf(w, "\n<br/>%v", c)
	}

	fmt.Fprintf(w, "\n<br/><h2>Configuration</h2><pre>%s</pre>",
		//strings.Replace(a.cnf.String(), "\n", "\n<br>", -1))
		a.cnf.String())

	// dump store.
	fmt.Fprintf(w, "\n<br/><h2>MemStore dump</h2><br/><pre>%v</pre>", a.str.String())

	fmt.Fprintf(w, "</html>")

}