package rolepermission

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
	store types.RolePermissionStore
}

func NewHandler(store types.RolePermissionStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/rolepermissions", h.handleGetRolePermissions).Methods(http.MethodGet)
	router.HandleFunc("/rolepermissions/{rolepermissionID}", h.handleGetRolePermissions).Methods(http.MethodGet)

}

func (h *Handler) handleGetRolePermissions(w http.ResponseWriter, r *http.Request) {
	rolepermissions, err := h.store.GetRolePermissions()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, rolepermissions)
}

func (h *Handler) handleGetRolePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["rolepermissionID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing rolepermission ID"))
		return
	}

	rolepermissionID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid rolepermission ID"))
		return
	}

	rolepermission, err := h.store.GetRolePermissionById(rolepermissionID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, rolepermission)
}

func (h *Handler) handleCreateRolePermission(w http.ResponseWriter, r *http.Request) {
	var rolepermission types.CreateRolePermissionPayload

	if err := utils.ParseJSON(r, &rolepermission); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	if err := utils.Validate.Struct(rolepermission); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateRolePermission(rolepermission)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, rolepermission)
}
