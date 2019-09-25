package common

// Mailer interface to sened emails.
type Mailer interface {
	Pinger
	Send(from, to, obj, txt string) error
}
