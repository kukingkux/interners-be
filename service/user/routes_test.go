package user

import (
	"fmt"

	"github.com/kukingkux/interners-be/types"
)

// func TestUserServiceHandlers(t *testing.T) {
// 	userStore := &mockUserStore{}
// 	handler := NewHandler(userStore)

// 	t.Run("should fail id user is invalid", func(t *testing.T) {
// 		payload := types.RegisterUserPayload{
// 			FirstName: "user",
// 			LastName:  "123",
// 			Email:     "aswan",
// 			Password:  "aswandala",
// 		}
// 		marshaled, _ := json.Marshal(payload)

// 		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshaled))
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		router := mux.NewRouter()

// 		router.HandleFunc("/register", handler.handleRegister)
// 		router.ServeHTTP(rr, req)

// 		if rr.Code != http.StatusBadRequest {
// 			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
// 		}
// 	})

// 	t.Run("should currently register the user", func(t *testing.T) {
// 		payload := types.RegisterUserPayload{
// 			FirstName: "user",
// 			LastName:  "123",
// 			Email:     "aswandala@valid.com",
// 			Password:  "aswandala",
// 		}
// 		marshaled, _ := json.Marshal(payload)

// 		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshaled))
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		router := mux.NewRouter()

// 		router.HandleFunc("/register", handler.handleRegister)
// 		router.ServeHTTP(rr, req)

// 		if rr.Code != http.StatusCreated {
// 			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
// 		}
// 	})
// }

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}
func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
func (m *mockUserStore) GetUsers(*types.User) error {
	return nil
}
