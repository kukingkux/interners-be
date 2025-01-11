package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "github.com/kukingkux/interners-be/config"

	"github.com/kukingkux/interners-be/types"
	"github.com/kukingkux/interners-be/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.handleGetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", h.handleUpdateUserAtFirstLogin).Methods(http.MethodPut)
}

func (h *Handler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) handleUpdateUserAtFirstLogin(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))
		return
	}
	
	err := h.store.UpdateUserAtFirstLogin(user)
	if err != nil {
		log.Printf("Error updating user: %v", err) 
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}