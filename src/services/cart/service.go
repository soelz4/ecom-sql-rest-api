package cart

import (
	"fmt"

	"ecom/src/types"
)

func GetCartItemsIDs(items []types.CartCheckoutItem) ([]int, error) {
	productIDs := make([]int, len(items))

	for i, v := range items {
		if v.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product %d", v.ProductID)
		}

		productIDs[i] = v.ProductID
	}

	return productIDs, nil
}

func checkIfCartIsInStock(
	cartItems []types.CartCheckoutItem,
	products map[int]types.Product,
) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]

		if !ok {
			return fmt.Errorf(
				"product %d is not available in the store, please refresh your cart",
				item.ProductID,
			)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(
	cartItems []types.CartCheckoutItem,
	products map[int]types.Product,
) float64 {
	var totalPrice float64 = 0

	for _, item := range cartItems {
		product := products[item.ProductID]
		totalPrice = totalPrice + (product.Price * float64(product.Quantity))
	}

	return totalPrice
}

// Return Order ID - Total Amount for the User to Pay - Error
func (h *Handler) createOrder(
	ps []types.Product,
	cartItems []types.CartCheckoutItem,
	userID int,
) (int, float64, error) {
	// Create a Map of Products for Easier Access
	productsMap := make(map[int]types.Product)
	for _, product := range ps {
		productsMap[product.ID] = product
	}

	// Check if All Products are Available
	err := checkIfCartIsInStock(cartItems, productsMap)
	if err != nil {
		return 0, 0, err
	}

	// Calculate Total Price
	totalPrice := calculateTotalPrice(cartItems, productsMap)

	// Reduce the Quantity of Products in the Store
	for _, item := range cartItems {
		product := productsMap[item.ProductID]
		product.Quantity = product.Quantity - item.Quantity
		h.productStore.UpdateProduct(product)
	}

	// Create Order Record
	orderID, err := h.orderStore.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "Pending",
		Address: "Some Address - IDK",
	})
	if err != nil {
		return 0, 0, err
	}

	// Create Order the Items Records
	for _, item := range cartItems {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productsMap[item.ProductID].Price,
		})
	}

	// Return
	return orderID, totalPrice, nil
}
