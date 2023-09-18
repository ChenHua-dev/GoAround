package main 

import (
	"fmt"
	"log"
    "net/http" 

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

func main() {
	// fmt.Println("hello world")

	// use net/http: 
	// fmt.Println("started-service")
    // http.HandleFunc("/upload", uploadHandler)
    // log.Fatal(http.ListenAndServe(":8080", nil))

	// use gorilla/mux: 
	fmt.Println("started-service")

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            return []byte(mySigningKey), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })
	
    r := mux.NewRouter()
    // r.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST", "OPTIONS")
	// r.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET", "OPTIONS")
	r.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST", "OPTIONS")
    r.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET", "OPTIONS")
	r.Handle("/post/{id}", jwtMiddleware.Handler(http.HandlerFunc(deleteHandler))).Methods("DELETE", "OPTIONS")
	r.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST", "OPTIONS")
    r.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST", "OPTIONS")

    log.Fatal(http.ListenAndServe(":8080", r))

}
