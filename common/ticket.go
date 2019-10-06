package common

import "time"

// Ticket data structure.
type Ticket struct {
	TID    string // Ticket ID
	Holder string // Ticket holder information
	Mail   string // Email to forward the ticket when issued
	Issued bool   // If not issued, ticket is not yet valid.
	From,
	To time.Time // Validity period
	Entries,
	Exits int // Counters
}

// Valid tests ticket validity.
// Adjust condition as needed.
func (t *Ticket) Valid() bool {
	return time.Now().After(t.From) &&
		time.Now().Before(t.To) &&
		t.Issued &&
		(t.Entries-t.Exits) <= 1

}
