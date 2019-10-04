package app

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/xavier268/go-ticket/common"
)

// tktHdlf display ticket publicly.
// Typically, used for printing before the event.
func (a *App) tktHdlf(w http.ResponseWriter, r *http.Request) {

	// Get deviceID and role

	ss := a.Authorize(w, r, common.RoleNone)

	content, err := a.str.Process(ss.TicketID, ss.Role)

	if err != nil {
		fmt.Fprintf(w, "<html><h1>This ticket does not exists in our base</h1>\n<br/>%s</html>", err.Error())
		return
	}

	u := path.Join(a.cnf.Addr.Public, a.cnf.API.QRImage)
	u = strings.Replace(u, "http:/", "http://", 1)
	u = strings.Replace(u, "https:/", "https://", -1)
	u = u + "?" + a.cnf.API.QueryParam.QRText + "=" + url.QueryEscape(ss.TicketID)

	fmt.Fprintf(w, "<html><h1>Ticket N° %s</h1>Information : %s<br/><img src='%s'><br/>You may print this ticket if you wish.</img></html>",
		ss.TicketID, content, u)

}
