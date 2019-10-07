package app

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/xavier268/go-ticket/common"
)

// MailTicketHdlf will mail a ticket or an activation link to the provided email.
// Mail and tickt number are passed as query parameters.
func (a *App) mailTicketHdlf(w http.ResponseWriter, r *http.Request) {

	// Only admin+ are authorized !
	ss := a.Authorize(w, r, common.RoleAdmin)

	if ss == nil {
		return
	}

	if len(ss.TicketID) == 0 || len(ss.Mail) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	txt := "A ticket has been prepared for you, please visit : "
	txt += a.CreateAbsoluteURL(a.cnf.API.Ticket) +
		"?" + a.cnf.API.QueryParam.Ticket + "=" +
		url.QueryEscape(ss.TicketID)

	html := a.textToHTML(txt)

	go a.ml.Send(a.cnf.MailSender, ss.Mail, "Your ticket ...", txt, html)
	fmt.Fprint(w, "OK, mail sent to ", ss.Mail)
	return

}

// Convert text to html using 'mail.html' template.
func (a *App) textToHTML(txt string) string {
	b := new(bytes.Buffer)
	a.cnf.ExecuteTemplate(b, "mail.html", txt)
	return b.String()
}
