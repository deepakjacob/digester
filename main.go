package main

import (
	//...

	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to digester"))
	})

	// Mount the digest sub-router
	r.Mount("/digest", digesterRouter())

	//TODO: implement graceful shutdown for server
	// go-chi reccomends following method:
	// github.com/go-chi/chi/blob/master/_examples/graceful/main.go
	http.ListenAndServe(":3333", r)
}

// A completely separate router for administrator routes
func digesterRouter() http.Handler {
	r := chi.NewRouter()
	// TODO: currently only basic auth with hardcoded username / password
	// TODO: change to stronger methods / implementations
	r.Use(AuthenticatedOnly)
	r.Get("/status", getStatus)
	r.Post("/post", postFile)
	return r
}

func _requestBasicAuth(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", `Basic realm="digester"`)
	w.WriteHeader(http.StatusUnauthorized)
}

// AuthenticatedOnly middleware restricts access to just authenticated requests only.
func AuthenticatedOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		isAuthenticated := username == "anil" && password == "passw0rd"
		if !ok || !isAuthenticated {
			_requestBasicAuth(w)
		}
		next.ServeHTTP(w, r)
	})
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello this is /getStatus"))
}

func postFile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello this is /postFile"))
}
