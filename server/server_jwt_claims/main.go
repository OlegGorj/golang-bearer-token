package main

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Define the secret key for signing and verifying tokens
var secret = []byte("my_secret_key")

// Define a custom Claims type that includes user information
type Claims struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Admin  bool   `json:"admin"`
	jwt.StandardClaims
}

// Define an HTTP endpoint for user login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the user credentials (e.g. by checking them against a database)
	if user.Email != "jane.doe@example.com" || user.Password != "password123" {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token with claims
	token, err := generateToken(user.ID, user.Name, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the token in the response body
	response := AuthResponse{Token: token}
	json.NewEncoder(w).Encode(response)
}

// Define an HTTP endpoint that requires authentication
func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the Authorization header to extract the token
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Use the claims to perform an action or return data
	w.Write([]byte("Hello, " + claims.Name + "!\n"))
	w.Write([]byte("Your user ID is " + claims.UserID + ".\n"))
	if claims.Admin {
		w.Write([]byte("You have administrator privileges.\n"))
	}
}

// Generate a JWT token with claims
func generateToken(userID, name string, isAdmin bool) (string, error) {
	claims := &Claims{
		UserID: userID,
		Name:   name,
		Admin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func main() {
	// Define HTTP routes
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/data", dataHandler)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
