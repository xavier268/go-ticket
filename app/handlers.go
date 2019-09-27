package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/xavier268/go-ticket/common/key"
	"github.com/xavier268/go-ticket/impl/barcode"
)

// pingHdlf generates a ping page, displaying various informations.
func (a *App) pingHdlf(w http.ResponseWriter, r *http.Request) {

	// Get/Set device id
	did := a.getDeviceID(w, r)

	// send response ...
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><h1>Ping response</h1>")

	fmt.Fprintf(w, "\n<br/><h2>Request</h2> <br/>Url : %s<br/>Device id : %s",
		r.URL, did)

	fmt.Fprintf(w, "\n<h2>Headers</h2>")
	for k, v := range r.Header {
		fmt.Fprintf(w, "\n<br/>%v : %v", k, v)
	}

	fmt.Fprintf(w, "\n<br/><h2>Cookies</h2>")
	for _, c := range r.Cookies() {
		fmt.Fprintf(w, "\n<br/>%v", c)
	}

	fmt.Fprintf(w, "\n<br/><h2>Configuration</h2><br/>%s",
		strings.Replace(a.cnf.String(), "\n", "\n<br>", -1))

	fmt.Fprintf(w, "</html>")
}

// qrHdlf displays qr code for the value provided in the request.
// the request url form should be encoded as http(s)://...?qr=urlencodedvalue
func (a *App) qrHdlf(w http.ResponseWriter, r *http.Request) {

	// read the qr parameter
	p := r.URL.Query().Get("qr")
	if a.cnf.GetBool(key.VERBOSE) {
		fmt.Println("Serving qr code for qr = ", p)
	}
	w.Header().Set("Content-type", "image/png")
	w.WriteHeader(http.StatusOK)
	barcode.New().SetFormat(barcode.QR200x200H).Encode(w, p)

}
