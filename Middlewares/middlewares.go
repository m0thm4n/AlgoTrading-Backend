package Middlewares

import (
	"errors"
	"net/http"

	"github.com/m0thm4n/AlgoTrading-Backend/Auth"
	"github.com/m0thm4n/AlgoTrading-Backend/Responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := Auth.TokenValid(r)
		if err != nil {
			Responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		next(w, r)
	}
}