package permission

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
	store types.PermissionStore
}

func NewHandler(store types.PermissionStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/permissions", h.handleGetPermissions).Methods(http.MethodGet)
	router.HandleFunc("/permissions/{roleID}", h.handleGetPermissions).Methods(http.MethodGet)

}

func (h *Handler) handleGetPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.store.GetPermissions()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, permissions)
}

func (h *Handler) handleGetPermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["roleID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing permission ID"))
		return
	}

	roleID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid permission ID"))
		return
	}

	permission, err := h.store.GetPermissionById(roleID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, permission)
}

func (h *Handler) handleCreatePermission(w http.ResponseWriter, r *http.Request) {
	var permission types.CreatePermissionPayload

	if err := utils.ParseJSON(r, &permission); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	if err := utils.Validate.Struct(permission); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreatePermission(permission)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, permission)
}
