package product

import (
	"github.com/gorilla/mux"

	"ecom/src/services/auth"
	"ecom/src/types"
)

type Handler struct {
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{productStore: productStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products/", h.handleGetProducts).Methods("GET")
	router.HandleFunc("/products/{productID}", h.handleGetProductByID).Methods("GET")
	router.HandleFunc("/products/", auth.WithJWTAuth(h.handleCreateProduct, h.userStore)).
		Methods("POST")
}
