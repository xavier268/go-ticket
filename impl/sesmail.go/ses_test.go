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

	m := NewSESMail()
	e := m.Ping()
	if e != nil {
		t.Fatal(e)
	}

	e = m.Send("xavier@twiceagain.com", "xavier.gandillot@gmail.com", "test email", time.Now().String(), "")
	if e != nil {
		t.Fatal(e)
	}
}
