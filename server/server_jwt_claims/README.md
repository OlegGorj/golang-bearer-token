# Bearer token authentication with JWT claims

Here's a breakdown of how the code works:

First, we define the secret key for signing and verifying tokens:
```go
var secret = []byte("my_secret_key")
```

Then, we define a custom type `Claims` that includes the user's information (ID, name, and admin status) as well as a `StandardClaims` struct that includes an expiration time:
```go
type Claims struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Admin  bool   `json:"admin"`
	jwt.StandardClaims
}
```

The `main()` function starts a super simple HTTP server and defines two endpoints: `/login` and `/data` where we define an HTTP endpoint for user login using the `http.HandleFunc()` function.

When a client sends a `POST` request to this endpoint with a JSON payload containing the user's email and password, the server performs the following steps:

- Decodes the JSON payload to get the user's email and password.
- Validates the user's email and password (e.g. by checking them against a database).
- If the email and password are valid, generates a JWT token with claims that include the user's ID, name, and admin status using the `generateToken()` function.
- Signs the JWT token with the secret key and returns the signed token in the response body as a JSON payload. 
In other words, the token is signed with a secret key and returned to the client in the response body.

```go
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
```

The `/data` endpoint requires the client to send a `GET` request with an Authorization header that includes the JWT token generated during the login process. 
The server parses the token, validates it with the secret key, and extracts the user's claims. 
It then uses the claims to perform an action or return data to the client. In this example, it returns a personalized greeting and information about the user's ID and admin status.

The `generateToken()` function creates a new `Claims` struct with the user's information and an expiration time, and signs it with a secret key. 
It returns the signed token as a string.

The `Claims` struct is a custom type that includes the user's ID, name, admin status, and a StandardClaims struct that includes an expiration time.

The `loginHandler()` and `dataHandler()` functions are HTTP request handlers that are called when the client sends a request to the `/login` or `/data` endpoints, respectively. 
They use the `http.ResponseWriter` and `http.Request` types to handle the request and response data.
