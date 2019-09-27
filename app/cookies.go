package app

import (
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, name, value string, age time.Duration) {
	if len(name) == 0 || len(value) == 0 {
		return
	}
	c := &http.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: int(age.Seconds())}
	http.SetCookie(w, c)
}

func readCookie(r *http.Request, name string) string {
	if len(name) != 0 {
		c, err := r.Cookie(name)
		if err == nil {
			return c.Value
		}
	}
	return ""
}

func deleteCookie(w http.ResponseWriter, name string) {
	if len(name) == 0 {
		return
	}
	c := &http.Cookie{
		Name:   name,
		MaxAge: -1}
	http.SetCookie(w, c)
}
