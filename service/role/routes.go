package role

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
	store types.RoleStore
}

func NewHandler(store types.RoleStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/roles", h.handleGetRoles).Methods(http.MethodGet)
	router.HandleFunc("/roles/{roleID}", h.handleGetRoles).Methods(http.MethodGet)

}

func (h *Handler) handleGetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.store.GetRoles()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, roles)
}

func (h *Handler) handleGetRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["roleID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing role ID"))
		return
	}

	roleID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid role ID"))
		return
	}

	role, err := h.store.GetRoleById(roleID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, role)
}

func (h *Handler) handleCreateRole(w http.ResponseWriter, r *http.Request) {
	var role types.CreateRolePayload

	if err := utils.ParseJSON(r, &role); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	if err := utils.Validate.Struct(role); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateRole(role)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, role)
}
