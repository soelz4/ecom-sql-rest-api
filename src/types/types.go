package types

import "time"

// Types
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderID"`
	ProductID int       `json:"productID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type CartCheckoutItem struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

// Storages
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProducts() ([]*Product, error)
	GetProductByID(id int) (*Product, error)
	GetProductsByIDs(id []int) ([]Product, error)
	CreateProduct(CreateProductPayload) error
	UpdateProduct(Product) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

// PayLoads
type RegisterUserPayload struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname"  validate:"required"`
	Email     string `json:"email"     validate:"required,email"`
	Password  string `json:"password"  validate:"required,min=3,max=110"`
}

type LoginUserPayload struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateProductPayload struct {
	Name        string  `json:"name"        validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"       validate:"required"`
	Quantity    int     `json:"quantity"    validate:"required"`
}

type CartCheckoutPayload struct {
	Item []CartCheckoutItem `json:"items" validate:"required"`
}
