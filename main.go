package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// JWKS represents a JSON Web Key Set
var jwks = `{
  "keys": [
    {
      "kty": "RSA",
      "kid": "example-key",
      "use": "sig",
      "n": "your_modulus_here",
      "e": "AQAB"
    }
  ]
}`

func jwksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jwks))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/.well-known/jwks.json", jwksHandler).Methods("GET")

	fmt.Println("JWKS server listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}