package cookie

import (
	"errors"
	"net/http"

	// "git.fs.bnf.net/bnfinet/lasso/pkg/structs"
	"git.fs.bnf.net/bnfinet/lasso/pkg/cfg"
	"git.fs.bnf.net/bnfinet/lasso/pkg/domains"
	log "github.com/Sirupsen/logrus"
)

var defaulMaxAge = cfg.Cfg.JWT.MaxAge * 60

// SetCookie http
func SetCookie(w http.ResponseWriter, r *http.Request, val string) {
	setCookie(w, r, val, defaulMaxAge)
}

func setCookie(w http.ResponseWriter, r *http.Request, val string, maxAge int) {
	// foreach domain
	if maxAge == 0 {
		maxAge = defaulMaxAge
	}
	domain := domains.Matches(r.Host)
	// log.Debugf("cookie %s expires %d", cfg.Cfg.Cookie.Name, expires)
	http.SetCookie(w, &http.Cookie{
		Name:     cfg.Cfg.Cookie.Name,
		Value:    val,
		Path:     "/",
		Domain:   domain,
		MaxAge:   maxAge,
		Secure:   cfg.Cfg.Cookie.Secure,
		HttpOnly: cfg.Cfg.Cookie.HTTPOnly,
	})
}

// Cookie get the lasso jwt cookie
func Cookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(cfg.Cfg.Cookie.Name)
	if err != nil {
		return "", err
	}
	if cookie.Value == "" {
		return "", errors.New("Cookie token empty")
	}
	log.Debugf("cookie %s: %s", cfg.Cfg.Cookie.Name, cookie.Value)
	return cookie.Value, err
}

// ClearCookie get rid of the existing cookie
func ClearCookie(w http.ResponseWriter, r *http.Request) {
	setCookie(w, r, "delete", -1)
}
