package app

import (
	"fmt"
	"net/http"

	"github.com/xavier268/go-ticket/impl/barcode"
)

// qrHdlf displays qr code for the value provided in the request.
// the request url form should be encoded as http(s)://...?qr=urlencodedvalue
func (a *App) qrHdlf(w http.ResponseWriter, r *http.Request) {

	// read the qr parameter
	p := r.URL.Query().Get("qr")
	if a.cnf.Test.Verbose {
		fmt.Println("Serving qr code for qr = ", p)
	}
	w.Header().Set("Content-type", "image/png")
	w.WriteHeader(http.StatusOK)
	barcode.New().SetFormat(barcode.QR200x200H).Encode(w, p)

}
