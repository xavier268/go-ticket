// Package e2e provides end-2-end testing for the app.
package e2e

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestPing(t *testing.T) {
	a, c := launch()
	defer a.Close()

	p, s := get(a.CreateAbsoluteURL(c.API.Ping))

	if s != http.StatusOK {
		fmt.Println("Status code = ", s)
		t.Fatal("Unexpected status code")
	}

	if !strings.Contains(p, "Running on version") {
		fmt.Println(p)
		t.Fatal("Unexpected response to ping")
	}

}

func TestAuthorization(t *testing.T) {
	a, c := launch()
	defer a.Close()

	// Unauthorized access should fail ...
	p, s := get(a.CreateAbsoluteURL(c.API.Admin))
	if s != http.StatusUnauthorized {
		fmt.Println(s, p)
		t.Fatal("Should have retunred an Unauthorized (401) status code")

	}

	// authorize access should succeed ...
	p, s = getWithCredentials(a.CreateAbsoluteURL(c.API.Admin), c.Superuser.Name, c.Superuser.Password)
	if s != http.StatusOK {
		fmt.Println(s, p)
		t.Fatal("Access should have been granted ?!")
	}
}
