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

	mess := "Testing ... " + time.Now().String()
	e = m.Send("xavier@gandillot.com", "xavier.gandillot@gmail.com", "test email", mess, mess)
	if e != nil {
		t.Fatal(e)
	}
}
