package app

import (
	"fmt"
	"testing"
	"time"

	"github.com/xavier268/go-ticket/common"
	"github.com/xavier268/go-ticket/conf"
)

// gettestConfig provides a test configuration.
func getTestConfig() *conf.Conf {
	return conf.NewConf()
}

func TestDump(t *testing.T) {
	getTestConfig().Dump()
}

func TestApp(t *testing.T) {
	a := NewApp(getTestConfig())

	// Create a test ticket
	tt := *MakeTicket("123456")
	a.str.SaveTicket(tt)
	fmt.Println("Just created a test ticket : ", tt)

	go a.Run() // run server for 5 minutes, then close it ...
	if testing.Short() {
		time.Sleep(5 * time.Second)
	} else {

		time.Sleep(15 * time.Minute)
	}
	a.Close()
}

func MakeTicket(tid string) *common.Ticket {
	// Create a  ticket
	tkt := new(common.Ticket)
	tkt.TID = tid
	tkt.Issued = true
	tkt.Holder = "M. or Mrs. HOLDER"
	tkt.Mail = "do.not@use.com"
	tkt.To = time.Now().Add(5 * time.Hour)
	tkt.From = time.Now()
	return tkt
}
