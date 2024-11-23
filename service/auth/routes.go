package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kukingkux/interners-be/config"
	"github.com/kukingkux/interners-be/types"
	"golang.org/x/oauth2"
)

type Handler struct {
	userStore types.UserStore
}

func NewAuthHandler(userStore types.UserStore) *Handler {
	return &Handler{userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth/google/login", h.oauthGoogleLogin).Methods(http.MethodGet)
	// router.HandleFunc("/auth/google/oauth", h.oauthHandler)
	router.HandleFunc("/auth/google/callback", h.oauthGoogleCallback).Methods(http.MethodGet)
	// router.HandleFunc("/logout", h.getAuthCallbackFunction)
}

func (h *Handler) oauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := generateStateOauthCookie(w)
	u := config.GoogleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

// func (h *Handler) oauthHandler(w http.ResponseWriter, r *http.Request) {
// 	url :=
// }

func (h *Handler) oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {

	user, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "User Info %s\n", user)

	// var jsonResp types.GoAuth
	// code := r.URL.Query().Get("Code")
	// t, err := config.GoogleOauthConfig.Exchange(context.Background(), code)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// client := config.GoogleOauthConfig.Client(context.Background(), t)
	// resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	// if err = json.NewDecoder(resp.Body).Decode(&jsonResp); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	http.SetCookie(w, &http.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	oauthGoogleUrlAPI := config.GoogleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	token, err := config.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	client := config.GoogleOauthConfig.Client(context.Background(), token)

	response, err := client.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	// contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	// saveUser(contents)
	// saveToken(contents, token)
	return []byte(token.AccessToken), nil
}
