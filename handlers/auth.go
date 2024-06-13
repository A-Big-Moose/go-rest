package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/A-Big-Moose/go-rest/pkg/jwt"
)

var users = map[string]string{
	"andrew": "michael1234",
	"user2":  "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, err := jwt.GenerateJWT(creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the token in the response body
	json.NewEncoder(w).Encode(TokenResponse{Token: tokenString})
}
