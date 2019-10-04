package app

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/xavier268/go-ticket/common"
)

// Display admin page.
// Require at least Admin role to get in.
func (a *App) adminHdlf(w http.ResponseWriter, r *http.Request) {

	ss := a.Authorize(w, r, common.RoleAdmin)
	if ss == nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><h1>Admin</h1>Current role : %s</html>\n", ss.Role.String())

	for i := 0; i < int(ss.Role); i++ {
		rr := common.Role(i)
		rs := rr.String()
		fmt.Fprintf(w, "\n<h2>Activate %s</h2>", rs)
		actRq := a.CreateActivationRequestURL(rr)
		imURL := path.Join("/", a.cnf.API.QRImage) +
			"?" + a.cnf.API.QueryParam.QRText +
			"=" + url.QueryEscape(actRq)

		fmt.Fprintf(w, "\n<img src=%s></img><br/>%s<br/>", imURL, actRq)
	}

	// Ping facility
	ping := path.Join(a.cnf.Addr.Public, a.cnf.API.Ping)
	ping = strings.Replace(ping, "http:/", "http://", -1)
	ping = strings.Replace(ping, "https:/", "https://", -1)
	imURL := path.Join("/", a.cnf.API.QRImage) +
		"?" + a.cnf.API.QueryParam.QRText +
		"=" + url.QueryEscape(ping)
	fmt.Fprintf(w, "\n<h2>Ping</h2><img src=%s></img><br/>%s<br/>", imURL, ping)
}

// CreateActivationRequestURL generate a Activation request,
// and package it in the public url.
// For instance, actRq = "http://public.com:8080/activate?act=123456789"
func (a *App) CreateActivationRequestURL(rr common.Role) string {
	actRq := path.Join(a.cnf.Addr.Public, a.cnf.API.Activate) +
		"?" + a.cnf.API.QueryParam.ActivationRequestID +
		"=" + a.str.CreateRequestID(rr)
	// Adjustments needed because
	actRq = strings.Replace(actRq, "http:/", "http://", 1)
	actRq = strings.Replace(actRq, "https:/", "http://", 1)
	return actRq
}
