package app

import (
	"fmt"
	"net/http"
)

func (a *App) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.logger.OutputInfo(fmt.Sprintf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))
		next.ServeHTTP(w, r)
	})
}
