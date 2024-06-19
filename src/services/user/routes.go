package user

import (
	"github.com/gorilla/mux"

	"ecom/src/types"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login/", h.handleLogin).Methods("POST")
	router.HandleFunc("/register/", h.handleRegister).Methods("POST")
}
