# golang-bearer-token


Bearer token authentication for APIs is a way to ensure that only authorized users can access specific API endpoints. 
Here's a simple explanation of how the workflow works:

1. The user logs in to the API using their username and password. This part implements authentication workflow against some trusted backend e.g. AD, etc.
2. If the login is successful, the API generates a unique "bearer token" for the user and sends it back in the response.
3. The user stores this token on their device or in their browser.
4. Whenever the user wants to access a protected endpoint on the API, they include the bearer token in the request.
5. The API checks the token to make sure it's valid and belongs to an authorized user.
6. If the token is valid, the API returns the requested data or performs the requested action. If the token is invalid, the API returns an error message.

One advantage of using bearer token authentication for APIs is that it allows users to authenticate once and then access multiple endpoints without having to log in again. This can make the user experience smoother and reduce the load on the server.

Here are some common use cases for bearer token authentication in APIs, of course this is not an exhausted list of use cases by no means:

- Social media apps, where users have different levels of access to data depending on their role (e.g. a user vs. an administrator).
- E-commerce sites, where users need to authenticate to access their shopping cart, order history, and other personal information.
- SaaS (Software as a Service) platforms, where different customers may have access to different features or data depending on their subscription plan.
- Mobile apps, where bearer tokens can be used to authenticate API requests even when the user is offline.

One thing to keep in mind is that bearer tokens can be vulnerable to attacks like interception or theft if they are not implemented securely. 
It's important to use HTTPS encryption to protect the token during transmission and store the token securely on the client side.

And, to quickly sum up pros and cons of using Bearer authentication method..
The pros:

- Bearer tokens are easy to implement and use. You just need to include the token in the Authorization header of your HTTP requests, and you're good to go. No need to worry about sessions or cookies.
- Bearer tokens are great for stateless environments like RESTful APIs, where you don't want to keep track of session information on the server. This makes them perfect for scaling out to handle a large number of clients.
- Bearer tokens are good for long-lived sessions. Since the token is stored on the client side, the server doesn't need to maintain any state. This means that the user can keep the same token for a long time without having to re-authenticate.
- Bearer tokens are versatile. They can be used in a variety of scenarios, such as mobile apps, web apps, and APIs.

Now, for the cons:

- Bearer tokens can be vulnerable to attacks like interception, replay, and theft if not implemented securely. It's important to use HTTPS encryption to protect the token during transmission, and to store the token securely on the client side.
- Bearer tokens can't be revoked. Once a token is issued, it's valid until it expires. You could set very short life-span for it 
but if you want to revoke a token, you'll need to wait for it to expire or change the secret key that's used to sign the tokens. 
This can be a bit of a hassle.
- Bearer tokens can be used by anyone who has access to them. 
If a token is stolen, an attacker can use it to impersonate the user and perform actions on their behalf. 
This is why it's important to keep the token secure and to use short expiration times to limit the window of vulnerability.

Having said all the above, my personal view that Bearer authentication should be used in fairly simple workflows that does not require presence of Central Authority to expire token on-demand.

### Bearer token with claims

One more point to mention - Bearer token authentication can be used with claims to provide additional security and functionality. This is very convent for some use cases.

In a nutshell, **claims** are pieces of information that are attached to a bearer token to provide additional context about the user or the application. 
For instance, you might include a user's role or permissions as a claim, which would allow you to restrict access to certain API endpoints based on the user's role (classics: User vs Admin).

To include claims in a bearer token, you typically encode them as a JSON Web Token (JWT). A JWT is a compact, URL-safe way of representing claims as a JSON object. 
The JWT includes the claims in the payload, along with a signature to ensure the integrity of the token.
A simple application using JWT auth flow: https://github.com/OlegGorj/golang-jwt-token


## A super simple app to implement Bearer token auth flow.

### Server side

Quick server side example, we define a few routes for our server:

We define a `User` struct to store user information, a `Token` struct to store the generated token value, and two global variables: `users` to hold the registered users, and `tokens` to store the generated tokens.
We then define two endpoints: `/login` to authenticate users and generate tokens, and `/data` to handle authenticated requests.

The `loginHandler` function handles the /login endpoint. It checks the user's credentials using the Basic Authentication scheme, and if they are valid, generates a new token and stores it in the tokens map. The function then returns a JSON response with the generated token value.

The `generateToken` function is a simple function that generates a new token. In this example, we simply hardcode the token value, but in a real application, you would use a more secure method for generating tokens.

The `authMiddleware` function is similar to the previous example. It takes a `http.HandlerFunc` and returns a new `http.HandlerFunc` that includes authentication checking logic. The middleware function retrieves the Bearer token from the request's Authorization header and checks if it is a valid token in the tokens map. If the token is missing or invalid, the middleware returns an error response to the client. Otherwise, it calls the next handler function passed as a parameter.

The `dataHandler` function is similar to the previous example, but it is only called if the token is validated by the middleware.

And, finally, we register the `/login` and `/data` endpoints with the `http.HandleFunc` function and start the server using `http.ListenAndServe`.

To compile binary:

```bash
cd server
make init   # init mod & deps
make build  # build binary
```

### Client side

Well, we first define a token variable with the bearer token value.

We then create an HTTP client using the `http.Client` struct. We use the `http.NewRequest` function to create an HTTP GET request to the URL `https://api.example.com/data` with an empty body. 
If there are any errors during the creation of the request, we use the panic function to stop execution.

Next, we set the `Authorization` header of the request to include the bearer token. We do this using the `req.Header.Set` function and passing in a string that includes the token value. 
The `fmt.Sprintf` function is used to format the string with the Bearer authentication scheme.
Here's a good read about Bearer authentication scheme: https://swagger.io/docs/specification/authentication/bearer-authentication/

We then execute the request using the HTTP client's `Do` method, which returns a response and an error. We use defer `resp.Body.Close()` to ensure that the response body is closed after it's been read.

Finally, we can process the response as needed. For example, we might use `ioutil.ReadAll` to read the response body into a byte slice and then convert it to a string using string(body).

To compile client binary:

```bash
cd client
make init   # init mod & deps
make build  # build binary
```

To sum it up, bearer token authentication is a flexible and secure way to control access to API endpoints and ensure that only authorized users can perform certain actions or access certain data. And, most importantly, its quite easy to implement and maintain.



