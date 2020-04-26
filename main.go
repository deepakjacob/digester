package main

import (
	//...

	"context"
	"flag"
	"net/http"
	"time"

	"github.com/deepakjacob/digester/db"
	"github.com/deepakjacob/digester/domain"
	"github.com/deepakjacob/digester/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	debug := flag.Bool("debug", true, "sets log level to debug")
	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

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

	log.Info().Msg("Server is ready to service requests.")
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
	w.Write([]byte("Hello this is /postFile"))
	conn := db.New(context.Background())
	db := &db.RegistrationDBImpl{
		PgConn: conn,
	}

	service := &service.RegistrationServiceImpl{
		RegistrationDB: db,
	}
	o := &domain.Registration{
		FileID:   "001",
		FileName: "Sample File",
	}
	service.RegisterFile(context.Background(), o)
}

func postFile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello this is /postFile"))

}
