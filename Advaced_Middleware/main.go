package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc // define a new middleware struct

// logging logs all requests with its path and the time it took to process
func Logging() Middleware {
	// create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// do middleware things
			start := time.Now()

			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			// call th next middleware/handler in chain
			f(w, r)
		}
	}
}

// method ensures that url can only be requested with a specific method, else return a 400 bad request
func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(w, r)
		}
	}
}

// chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	fmt.Println(" * run server on 9090 port")
	http.ListenAndServe(":9090", nil)
}
