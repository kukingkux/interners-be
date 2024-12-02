package company

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
	store types.CompanyStore
}

func NewHandler(store types.CompanyStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/companies", h.handleGetCompanies).Methods(http.MethodGet)
	router.HandleFunc("/companies/{companyID}", h.handleGetCompanies).Methods(http.MethodGet)

}

func (h *Handler) handleGetCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := h.store.GetCompanies()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, companies)
}

func (h *Handler) handleGetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["companyID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing company ID"))
		return
	}

	companyID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid company ID"))
		return
	}

	company, err := h.store.GetCompanyById(companyID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, company)
}

func (h *Handler) handleCreateCompany(w http.ResponseWriter, r *http.Request) {
	var company types.CreateCompanyPayload

	if err := utils.ParseJSON(r, &company); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	if err := utils.Validate.Struct(company); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateCompany(company)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, company)
}
