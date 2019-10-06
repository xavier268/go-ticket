package app

import (
	"os"
	"testing"

	"github.com/xavier268/go-ticket/common"
	"github.com/xavier268/go-ticket/conf"
)

func TestTemplates(t *testing.T) {

	// Create test session object
	ss := new(SessionData)
	ss.Conf = conf.NewConf()
	ss.Writer = os.Stdout
	ss.Role = common.RoleAdmin

	// call templates
	ss.ExecuteTemplate("main.html")

}
