# Bearer token authentication with JWT claims


Here's a breakdown of how the code works:

The `main()` function starts a suoer simple HTTP server and defines two endpoints: `/login` and `/data`.

The `/login` endpoint handles a `POST` request from the client. 
It expects a JSON payload with the user's email and password, which it validates against a hardcoded value. 
If the email and password are valid, the server generates a JWT token with claims that include the user's ID, name, and admin status. 
The token is signed with a secret key and returned to the client in the response body.

The `/data` endpoint requires the client to send a `GET` request with an Authorization header that includes the JWT token generated during the login process. 
The server parses the token, validates it with the secret key, and extracts the user's claims. 
It then uses the claims to perform an action or return data to the client. In this example, it returns a personalized greeting and information about the user's ID and admin status.

The `generateToken()` function creates a new `Claims` struct with the user's information and an expiration time, and signs it with a secret key. 
It returns the signed token as a string.

The `Claims` struct is a custom type that includes the user's ID, name, admin status, and a StandardClaims struct that includes an expiration time.

The `loginHandler()` and `dataHandler()` functions are HTTP request handlers that are called when the client sends a request to the `/login` or `/data` endpoints, respectively. 
They use the `http.ResponseWriter` and `http.Request` types to handle the request and response data.
