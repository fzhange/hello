package demo

import (
	"fmt"
	"hello/demo/token"

	"github.com/beego/beego/logs"
)

/**
What the heck is a JWT?
In short, it's a signed JSON object that does something useful (for example, authentication).
It's commonly used for Bearer tokens in Oauth 2.
A token is made of three parts, separated by .'s.

The first two parts are JSON objects, that have been base64url encoded.
The last part is the signature, encoded the same way.

The first part is called the header.
	It contains the necessary information for verifying the last part, the signature.
	For example, which encryption method was used for signing and what key was used.

The part in the middle is the interesting bit.
	It's called the Claims and contains the actual stuff you care about.
	Refer to RFC 7519 for information about reserved keys and the proper way to add your own.

The last part is signature for first and second part.
**/

/**
For example, when a user logs in to a website secured via JWTs, the flow should look something like this:
1. The user sends a username and password to the server
2. The server verifies username and password are correct
3. The server creates a JSON object (also known as the “claims”) that looks something like this: {"username":"wagslane"}
4. The server encodes and signs the JSON object, creating a JWT:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IndhZ3NsYW5lIn0.ov6d8XtwQoKUwsYahk9UwH333NICElFSs6ag6pINyPQ
5. The user’s web client saves the JWT for later use
6. When the user makes a request to a protected endpoint, it passes the JWT along in an HTTP header
7. The server checks the signature on the JWT to make sure the JWT was originally created by the same server
8. The server reads the claims and gives permission to the request to operate as “wagslane”
**/

func Initiator() {
	var secretKey = "12345678901234567890123456789012"
	jwtMaker, err := token.NewJWTMaker(secretKey)
	if err != nil {
		logs.Error(err)
		// return
	}
	tokenStr, err := jwtMaker.CreateToken("fzhange", 8888000)
	if err != nil {
		logs.Error(err)
		// return
	}
	logs.Info("tokenStr", tokenStr)
	// var expiredTokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjFmZjFmMTY1LTY0NmEtNDY2Ni04OWU1LTQ1ZWM2MjZiMmFmZiIsInVzZXJuYW1lIjoiZnpoYW5nZSIsImlzc3VlZF9hdCI6IjIwMjItMTAtMjdUMTk6NTk6MjEuNDQ3KzA4OjAwIiwiZXhwaXJlZF9hdCI6IjIwMjItMTAtMjdUMTk6NTk6MjEuNDQ3MDAxKzA4OjAwIn0.-4UvbNCxD3BgTEQWTnQHj7unzrKz-0H4-BV9Jzjcneo"
	paylaod, err := jwtMaker.VerifyToken(tokenStr)
	fmt.Printf("%+v", paylaod, err)
}
