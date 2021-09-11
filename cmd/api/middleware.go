package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
			next.ServeHTTP(w, r)
		})
}

func (app *application) checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// could set anonymous user
		}
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			app.errorJSON(w, errors.New("invalid auth header"))
			return
		}
		if headerParts[0] != "Bearer" {
			app.errorJSON(w, errors.New("unautherized - no bearer"))
			return
		}

		token := headerParts[1]
		claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secret))
		if err != nil {
			app.errorJSON(w, errors.New("unautherized - HMAC check invalid"), http.StatusForbidden)
			return
		}
		if !claims.Valid(time.Now()) {
			app.errorJSON(w, errors.New("unautherized - token expired"), http.StatusForbidden)
			return
		}
		if !claims.AcceptAudience("mydomain.com") {
			app.errorJSON(w, errors.New("unautherized - invalid audience"), http.StatusForbidden)
			return
		}
		if claims.Issuer != "mydomain.com" {
			app.errorJSON(w, errors.New("unautherized - invalid issuer"), http.StatusForbidden)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			app.errorJSON(w, errors.New("unautherized - unautherized"), http.StatusForbidden)
			return
		}

		log.Println("Valid User: ", userID)
		next.ServeHTTP(w, r)
	})
}
