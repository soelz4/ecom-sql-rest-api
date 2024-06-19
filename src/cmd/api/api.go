package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"ecom/src/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	// Router and SubRouter
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// User Handler
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	// Start Server at PORT 9010
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
