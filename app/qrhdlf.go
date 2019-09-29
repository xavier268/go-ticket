package app

import (
	"fmt"
	"net/http"
)

// qrHdlf displays qr code for the value provided in the request.
// the request text should be encoded as query parameter QRText
func (a *App) qrHdlf(w http.ResponseWriter, r *http.Request) {

	// read the qr parameter
	p := r.URL.Query().Get(a.cnf.API.QueryParam.QRText)
	if a.cnf.Test.Verbose {
		fmt.Println("Serving qr code for qr = ", p)
	}
	w.Header().Set("Content-type", "image/png")
	w.WriteHeader(http.StatusOK)
	a.bc.Encode(w, p)

}
