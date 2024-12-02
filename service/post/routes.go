package post

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
	store types.PostStore
}

func NewHandler(store types.PostStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/posts", h.handleGetPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts/{postID}", h.handleGetPost).Methods(http.MethodGet)

	router.HandleFunc("/posts", h.handleCreatePost).Methods(http.MethodPost)
	router.HandleFunc("/posts", h.handleUpdatePost).Methods(http.MethodPut)
	router.HandleFunc("/posts", h.handleDeletePost).Methods(http.MethodDelete)
}

func (h *Handler) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.store.GetPosts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}

func (h *Handler) handleGetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["postID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing post ID"))
		return
	}

	postID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid post ID"))
		return
	}

	post, err := h.store.GetPostById(postID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, post)
}

func (h *Handler) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	var post types.CreatePostPayload

	if err := utils.ParseJSON(r, &post); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(post); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.CreatePost(post)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, post)
}

func (h *Handler) handleUpdatePost(w http.ResponseWriter, r *http.Request) {
	var post types.Post

	if err := utils.ParseJSON(r, &post); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.UpdatePost(post)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, post)
}

func (h *Handler) handleDeletePost(w http.ResponseWriter, r *http.Request) {
	var post types.Post

	if err := utils.ParseJSON(r, &post); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.DeletePost(post)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, post)
}
