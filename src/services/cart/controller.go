package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"ecom/src/services/auth"
	"ecom/src/types"
	"ecom/src/utils"
)

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	var cart types.CartCheckoutPayload
	userID := auth.GetUserIDFromContext(r.Context())

	err := utils.ParseJSON(r, &cart)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.Validate.Struct(cart)
	if err != nil {
		err = err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err))
		return
	}

	productIDs, err := GetCartItemsIDs(cart.Item)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// GET Products
	ps, err := h.productStore.GetProductsByIDs(productIDs)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Order ID - Total Price
	orderID, totalPrice, err := h.createOrder(ps, cart.Item, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"Order ID":    orderID,
		"Total Price": totalPrice,
	})
}
