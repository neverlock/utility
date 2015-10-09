package httplog

import (
	"html"
	"log"
	"net/http"
)

func HttpLogln(r *http.Request) {
	log.Printf("%s - %s - %s - %s - %q - %s",
		r.RemoteAddr,
		r.Proto,
		r.Method,
		r.UserAgent(),
		html.EscapeString(r.URL.Path),
		r.RequestURI,
	)
}
