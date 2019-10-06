package app

import (
	"net/http"
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

	ss.ExecuteTemplate("header.html")

	type param struct {
		U string // url to be encoded in the qr code
		T string // Text to display
	}

	for i := 0; i < int(ss.Role); i++ {
		data := new(param)
		r := common.Role(i)
		data.U = a.CreateActivationRequestURL(r)
		data.T = "Scan me to login for " + r.String()
		a.cnf.ExecuteTemplate(w, "adminFragment.html", data)
	}

	// Ping facility
	data := new(param)
	data.U = a.CreateAbsoluteURL(a.cnf.API.Ping)
	data.T = "Scan me to ping the server"
	a.cnf.ExecuteTemplate(w, "adminFragment.html", data)

	// Logout
	data = new(param)
	data.U = a.CreateAbsoluteURL(a.cnf.API.Logout)
	data.T = "Scan me to logout"
	a.cnf.ExecuteTemplate(w, "adminFragment.html", data)

	ss.ExecuteTemplate("footer.html")

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

// CreateAbsoluteURL from a relative url using public address.
func (a *App) CreateAbsoluteURL(relativeURL string) string {
	u := path.Join(a.cnf.Addr.Public, relativeURL)
	u = strings.Replace(u, "http:/", "http://", 1)
	u = strings.Replace(u, "https:/", "https://", 1)
	return u

}
