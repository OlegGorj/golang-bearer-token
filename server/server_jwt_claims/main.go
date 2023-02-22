package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Value     string    `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
}

type APIError struct {
	Message string `json:"message"`
}

const tokenDuration = time.Hour
const authToken = "my_bearer_token" // Replace with your actual token

var tokenMap map[string]Token = make(map[string]Token)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/token", tokenHandler)
	http.HandleFunc("/data", authMiddleware(dataHandler))
	http.ListenAndServe(":8080", nil)
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, exists := tokenMap[tokenString]
		if !exists {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if token.ExpiresAt.Before(time.Now()) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			delete(tokenMap, tokenString)
			return
		}

		next(w, r)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// this part implements authentication workflow
	// 	check if the user is valid (e.g., check against a AD)
	if user.Username != "example_user" || user.Password != "example_password" {
		errorResponse := APIError{Message: "Invalid username or password"}
		jsonResponse, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonResponse)
		return
	}

	// Generate and store a new token
	tokenValue := fmt.Sprintf("Bearer %d", time.Now().Unix())
	token := Token{Value: tokenValue, ExpiresAt: time.Now().Add(tokenDuration)}
	tokenMap[tokenValue] = token

	// Return the token in the response
	jsonResponse, _ := json.Marshal(token)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	var token Token
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the token is valid and not expired
	storedToken, exists := tokenMap[token.Value]
	if !exists || storedToken.ExpiresAt.Before(time.Now()) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		delete(tokenMap, token.Value)
		return
	}

	jsonResponse, _ := json.Marshal(storedToken)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Process the request here
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		return
	}
}
