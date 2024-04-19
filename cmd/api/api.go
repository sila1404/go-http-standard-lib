package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/sila1404/go-http-standard-lib/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// NewAPIServer creates a new instance of APIServer.
//
// Parameters:
// - addr: the address for the server.
// - db: a pointer to a sql.DB instance.
// Return type: *APIServer
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the API server and listens for incoming HTTP requests.
//
// It creates a new HTTP server with a ServeMux router and sets up a versioned API
// endpoint at "/api/v1/". It registers the user service handler and starts
// listening on the specified address.
//
// Returns an error if the server fails to start.
func (s *APIServer) Run() error {
	router := http.NewServeMux()

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	userStore := user.NewStore(s.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoute(v1)

	log.Println("Listening and serving HTTP on ", s.addr)
	return http.ListenAndServe(s.addr, v1)
}
