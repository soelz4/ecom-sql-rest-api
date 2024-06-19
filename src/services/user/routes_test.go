package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"ecom/src/types"
)

type mockUserStore struct{}

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	// Fist Test Case - Validate Register User PayLoad
	t.Run("Should Fail if the User PayLoad is Invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "hesad",
			Email:     "invalid",
			Password:  "1234",
		}

		marshalled_payload, _ := json.Marshal(payload)

		req, err := http.NewRequest(
			http.MethodPost,
			"/register/",
			bytes.NewBuffer(marshalled_payload),
		)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register/", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code == http.StatusBadRequest {
			t.Errorf("expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
		}
	})

	// 2nd Test Case - Validate Register User
	t.Run("Should Correctly Register the User", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "hesad",
			Email:     "valid@mail.com",
			Password:  "1234",
		}

		marshalled_payload, _ := json.Marshal(payload)

		req, err := http.NewRequest(
			http.MethodPost,
			"/register/",
			bytes.NewBuffer(marshalled_payload),
		)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register/", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, fmt.Errorf("hech hesad")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
