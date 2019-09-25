//Package mockmail is used for debugging mailers.
// It pretend it is sending emails, by displaying mail content on stout.
package mockmail

import (
	"fmt"

	"github.com/xavier268/go-ticket/common"
)

// MockMailer simulates sending emails.
type MockMailer struct{}

// Compiler checks
var _ common.Mailer = new(MockMailer)

// New creates a new MockMailer.
func New() *MockMailer {
	return new(MockMailer)
}

// Ping for health.
func (m *MockMailer) Ping() error {
	return nil
}

// Send sends a mail.
func (m *MockMailer) Send(from, to, obj, txt string) error {
	fmt.Printf("\nFROM : %s\nTO   : %s\n,OBJ   : %s\n%s\n", from, to, obj, txt)
	return nil
}
