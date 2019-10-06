package app

import (
	"fmt"
	"net/http"

	"github.com/xavier268/go-ticket/common"
)

// tktHdlf display ticket publicly.
// Typically, used for printing before the event.
func (a *App) tktHdlf(w http.ResponseWriter, r *http.Request) {

	// Get deviceID and role

	ss := a.Authorize(w, r, common.RoleNone)
	if ss == nil {
		fmt.Println("Unexpected internal error")
		return
	}

	// Read ticket for session
	ss.Ticket = a.str.GetTicket(ss.TicketID)

	// Display current status
	ss.ExecuteTemplate("ticket.html")

	// Update ticket data
	if ss.Role == common.RoleEntry && ss.Ticket.Valid() {
		ss.Ticket.Entries++
	}
	if ss.Role == common.RoleExit {
		ss.Ticket.Exits++
	}

	a.str.SaveTicket(ss.Ticket)

}
