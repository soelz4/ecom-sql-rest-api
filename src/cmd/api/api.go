package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"ecom/src/services/cart"
	"ecom/src/services/order"
	"ecom/src/services/product"
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
	// Product Handler
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)
	// Order Store
	orderStore := order.NewStore(s.db)
	// Cart Handler
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)
	// Start Server at PORT 9010
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
