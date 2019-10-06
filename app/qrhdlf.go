package app

import (
	"fmt"
	"net/http"

	"github.com/xavier268/go-ticket/common"
)

// qrHdlf displays qr code for the value provided in the request.
// the request text should be encoded as query parameter QRText
func (a *App) qrHdlf(w http.ResponseWriter, r *http.Request) {

	ss := a.Authorize(w, r, common.RoleNone)

	if a.cnf.Test.Verbose {
		fmt.Println("Serving qr code for qr = ", ss.QRTxt)
	}
	w.Header().Set("Content-type", "image/png")
	w.WriteHeader(http.StatusOK)
	a.bc.Encode(w, ss.QRTxt)
}
