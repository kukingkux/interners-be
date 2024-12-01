package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kukingkux/interners-be/config"
	"github.com/kukingkux/interners-be/types"
	"golang.org/x/oauth2"
)

type Handler struct {
	userStore types.UserStore
	goAuth    types.GoAuth
	db        *sql.DB
	certs     map[string]string
}

func NewAuthHandler(userStore types.UserStore) (*Handler, error) {
	certs, err := Certificates()
	if err != nil {
		return nil, err
	}

	return &Handler{certs: certs, userStore: userStore}, nil

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

	assertion := r.Header.Get("X-Goog-IAP-JWT-Assertion")
	if assertion == "" {
		fmt.Fprintln(w, "No Cloud IAP header found.")
	}

	email, _, err := ValidateAssertion(assertion, h.certs)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "Could not validate assertion. Check app logs.")
		return
	}

	fmt.Fprintf(w, "Hello %s\n", email)

}

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
	// a, err := NewApp()
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	// var userData types.GoAuth
	// err = json.Unmarshal(contents, &userData)
	// if err != nil {
	// 	return nil, fmt.Errorf("JSON parsing failed: %s", err.Error())
	// }

	// var existingUser User
	// if err := db.Query("Wemail = ?", types.GoAuth.Email).First(&existingUser).Error; err != nil {

	// }

	// saveUser(contents)
	// saveToken(contents, token)
	return contents, nil
}
