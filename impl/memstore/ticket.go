package memstore

import (
	"fmt"
	"time"

	"github.com/xavier268/go-ticket/common"
)

// Process ticketID, with a given role.
// Return human readable feedback, error.
func (s *MemStore) Process(tktID string, role common.Role) (string, error) {

	tkt, ok := s.tkt[tktID]
	ok = ok && tkt.valid()

	switch role {

	case common.RoleExit: // exiting ...
		if ok {
			tkt.exit++
			s.tkt[tktID] = tkt
			return "Good bye " + tkt.ticketHolder, nil
		}
		return "Ticket is not valid", common.ErrorInvalidTicket

	case common.RoleEntry: // entering
		if ok {
			tkt.entries++
			s.tkt[tktID] = tkt
			return "Welcome " + tkt.ticketHolder, nil
		}
		return "Ticket is not valid", common.ErrorInvalidTicket

	case common.RoleReview, common.RoleAdmin, common.RoleSuper: // just looking ...
		if ok {
			return tkt.String(), nil
		}
		return tkt.String(), common.ErrorInvalidTicket
	default: // do nothing
		return "You are not authorized to process this ticket", nil
	}
}

// ===============================internal ticket struct================================

// Internal ticket structure.
type ticket struct {
	tid          string // ticket id
	from         time.Time
	to           time.Time
	paid         bool   // payment made ?
	ticketHolder string // name of holder
	entries      int    // entry counts
	exit         int    // exit count
}

// String ticket
func (t *ticket) String() string {
	return fmt.Sprintf("Ticket N° %s\nTicket holder : %s\nValid from : %v\n     to: %v\nPaid ? : %v\n Entries/Exits : %d / %d\n",
		t.tid, t.ticketHolder, t.from, t.to, t.paid, t.entries, t.exit)
}

// Internal validation process.
// Refuse out of time window, unpaid, and no entry yet.
func (t *ticket) valid() bool {
	return time.Now().Before(t.to) &&
		time.Now().After(t.from) &&
		t.paid &&
		t.entries <= 0
}
