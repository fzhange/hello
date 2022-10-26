package demo

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
import (
	"fmt"

	"github.com/beego/beego/logs"
	"github.com/golang-jwt/jwt/v4"
)

var signKey = []byte("SecretYouShouldHide")

func createToken() (string, error) {

	claims := &jwt.StandardClaims{
		ExpiresAt: 3600 * 1000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(signKey)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"foo":     "bar",
	// 	"nbf":     time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// 	"expired": time.Now().Add(1 * time.Minute),
	// })
	// return token.SignedString(signKey)
}

func validate(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	// expired := "expired"

	fmt.Println("You>>>", token.Claims)
	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}

func StartServer() {
	if tokenString, err := createToken(); err != nil {
		logs.Error(err)
	} else {
		// var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"
		fmt.Println(">>>", tokenString)
		// validate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoiMjAyMi0xMC0yNlQxOTozMzozOS45OTE4MzYrMDg6MDAiLCJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.Et6PQMbqiln6CzWo81ONtSeM0C7Ceu0z1MphpVMtE7k")
		validate(tokenString)
	}

	// Token from another example.  This token is expired
	// var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("AllYourBase"), nil
	// })

	// if token.Valid {
	// 	fmt.Println("You look nice today")
	// } else if ve, ok := err.(*jwt.ValidationError); ok {
	// 	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
	// 		fmt.Println("That's not even a token")
	// 	} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
	// 		// Token is either expired or not active yet
	// 		fmt.Println("Timing is everything")
	// 	} else {
	// 		fmt.Println("Couldn't handle this token:", err)
	// 	}
	// } else {
	// 	fmt.Println("Couldn't handle this token:", err)
	// }
}
