//Package sesmail implements mail sending using AWS SimpleEmailService (SES).
package sesmail

import (
	"errors"
	"fmt"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"github.com/xavier268/go-ticket/common"
)

// SesMail is the mail sending object.
type SesMail struct {
	svc *ses.SES
}

// Compiler checks SESMail is a Mailer.
var _ common.Mailer = new(SesMail)

// New creates a new mail sending objet.
// The correct AWS credentials must be available loacally or via the envirnement.
// (see AWS documentation)
func New() *SesMail {
	var err error

	m := new(SesMail)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		fmt.Println(err)
		panic("Could not initiate a SESMail Session ?!")
	} else {
		m.svc = ses.New(sess)
	}
	return m
}

// Ping check health status.
// Very basic at this stage, consider actually sending a test mail ?
func (m *SesMail) Ping() error {
	if m.svc != nil {
		return nil
	}
	return errors.New("SES Mailer could not be created")
}

// Send the provided email.
func (m *SesMail) Send(from, to, obj, txt string, html string) (err error) {

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(to)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(html),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(txt),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(obj),
			},
		},
		Source: aws.String(from),
	}

	_, err = m.svc.SendEmail(input)

	if err != nil {
		fmt.Println(err)
	}

	return err
}
