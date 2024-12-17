package userrole

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/kukingkux/interners-be/types"
	"github.com/kukingkux/interners-be/utils"
)

type Handler struct {
	store types.UserRoleStore
}

func NewHandler(store types.UserRoleStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/userroles", h.handleGetUserRoles).Methods(http.MethodGet)
	router.HandleFunc("/userroles/{userroleID}", h.handleGetUserRoles).Methods(http.MethodGet)

}

func (h *Handler) handleGetUserRoles(w http.ResponseWriter, r *http.Request) {
	userroles, err := h.store.GetUserRoles()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, userroles)
}

func (h *Handler) handleGetUserRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userroleID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing userrole ID"))
		return
	}

	userroleID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid userrole ID"))
		return
	}

	userrole, err := h.store.GetUserRoleById(userroleID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, userrole)
}

func (h *Handler) handleCreateUserRole(w http.ResponseWriter, r *http.Request) {
	var userrole types.CreateUserRolePayload

	if err := utils.ParseJSON(r, &userrole); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	if err := utils.Validate.Struct(userrole); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateUserRole(userrole)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, userrole)
}
