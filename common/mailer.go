package common

// Mailer interface to send emails.
type Mailer interface {
	Pinger
	Send(from, to, obj, txt, html string) error
}
