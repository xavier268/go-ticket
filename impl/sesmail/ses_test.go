package sesmail

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSendMail(t *testing.T) {

	if _, ok := os.LookupEnv("TRAVIS"); ok {
		fmt.Println("Skipping test emailing from Travis")
		t.Skip()
	}

	m := New()
	e := m.Ping()
	if e != nil {
		t.Fatal(e)
	}

	txt := "\n\nTesting ... " + time.Now().String()
	html := "<html><body><br/>" + txt + "</body></html>"
	e = m.Send("go-ticket@yopmail.com", "xavier@yopmail.com", "test email", txt, html)
	if e != nil {
		t.Fatal(e)
	}
}
