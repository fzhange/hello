package demo

/*
* reference article: https://blog.logrocket.com/jwt-authentication-go/
										https://blog.boot.dev/golang/jwts-in-golang/

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
*/

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// In production make sure you use a real private key, preferably at least 256 bits in length:
var sampleSecretKey = []byte("SecretYouShouldHide")

func generateJWT() (string, error) {
	type customClaims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		Username: "user-name",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000, //time.Now().Add(10 * time.Minute),
			Issuer:    "nameOfWebsiteHere",
		},
	})

	if signedToken, err := token.SignedString(sampleSecretKey); err != nil {
		log.Fatal(err)
		return "", err
	} else {
		return signedToken, nil
	}
}

func verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// str, err := generateJWT()

		if request.Header["Token"] != nil {

		}
		writer.Write([]byte("123121"))
	})
}

func StartServer() {
	log.Println("111")
	http.HandleFunc("/home", verifyJWT(handlePage))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("000")
		log.Fatal(err)
	} else {
		log.Println("server has started now")
	}
}

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

/*
* The handlePage handler function will return the encoded JSON of the Message struct as a response to the client
* if the request is authorized after the request body is encoded.
 */

func handlePage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(request.Body).Decode(&message)
	if err != nil {
		return
	}
	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		return
	}
}
