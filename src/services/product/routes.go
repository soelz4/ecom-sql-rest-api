package product

import (
	"github.com/gorilla/mux"

	"ecom/src/types"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products/", h.handleGetProducts).Methods("GET")
	router.HandleFunc("/products/{productID}", h.handleGetProductByID).Methods("GET")
	router.HandleFunc("/products/", h.handleCreateProduct).Methods("POST")
}
