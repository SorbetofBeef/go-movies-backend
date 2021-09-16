package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SorbetofBeef/go-movies-backend/models"
	"github.com/pascaldekloe/jwt"

	"golang.org/x/crypto/bcrypt"
)

var validUser = models.User{
	ID:       10,
	Email:    "me@here.com",
	Password: "$2a$12$k4z7htWADpuASTDKIZLo/.zppKPsPoi/AseEGaUI80lCweB1xHWUG",
	// Password: "$2a$12$93KF3JVuh2oQlzJQdid2IOXP67F3xhI6fwD9yQ9bxM.5TtuCP4CvS",
}

// Credentials is the type for user inputted credentials
type Credentials struct {
	UserName string `json:"email"`
	Password string `json:"password"`
}

func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}

	// "$2a$12$93KF3JVuh2oQlzJQdid2IOXP67F3xhI6fwD9yQ9bxM.5TtuCP4CvS"
	hashedPassword := validUser.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		app.errorJSON(w, errors.New("error signing"))
		return
	}
	app.writeJSON(w, http.StatusOK, string(jwtBytes), "response")
}
