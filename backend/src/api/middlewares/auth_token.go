package middlewares

import (
	"api"
	"interfaces"
	"net/http"
)

type AuthSession struct {
	Service interfaces.SessionAccessorService
}

func (a AuthSession) HasValidSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer api.SendErrorIfPanicked(w)

		_, err := a.Service.GetSession(r)
		if err == nil {
			next.ServeHTTP(w, r)
		} else {
			panic(NoSession)
		}
	})
}
