package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// Custom servermux is a struct which can be multiplexers
type CustomServermux struct {
}

// This is a function handler to be overridden
func (p *CustomServermux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is %f", rand.Float64())
}

func main() {
	// Any struct that has serverHttp function can be multiplexers
	mux := &CustomServermux{}
	http.ListenAndServe(":8080", mux)
}
