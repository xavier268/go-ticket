package app

import (
	"net/http"

	"github.com/xavier268/go-ticket/common"
)

// MailTicketHdlf will mail a ticket or an activation link to the provided email.
// Format is /.../mail_API/tkt_API/
// Mail and tickt number are passed as query parameters.
func (a *App) mailTicketHdlf(w http.ResponseWriter, r *http.Request) {

	// Only admin+ are authorized !
	ss := a.Authorize(w, r, common.RoleAdmin)

	if ss == nil || len(ss.TicketID) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
