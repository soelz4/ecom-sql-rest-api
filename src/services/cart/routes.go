package cart

import (
	"github.com/gorilla/mux"

	"ecom/src/services/auth"
	"ecom/src/types"
)

type Handler struct {
	orderStore   types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(
	orderStore types.OrderStore,
	productStore types.ProductStore,
	userStore types.UserStore,
) *Handler {
	return &Handler{orderStore: orderStore, productStore: productStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout/", auth.WithJWTAuth(h.handleCheckout, h.userStore)).
		Methods("POST")
}
