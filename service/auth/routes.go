package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/kukingkux/interners-be/config"
	"github.com/kukingkux/interners-be/types"
	"github.com/kukingkux/interners-be/utils"
	"golang.org/x/oauth2"
)

type Handler struct {
	userInfo types.UserInfo
	userStore types.UserStore
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
	router.HandleFunc("/auth/google/callback", h.oauthGoogleCallback).Methods(http.MethodGet)
	router.HandleFunc("/auth/user", h.handleGetUser).Methods(http.MethodGet)
	// router.HandleFunc("/auth/user", h.handleCreateUser).Methods(http.MethodPost)
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
	// Retrieve the state from cookies and compare it with the one received in the callback
	oauthState, err := r.Cookie("oauthstate")
	if err != nil || r.FormValue("state") != oauthState.Value {
		log.Println("Invalid OAuth state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Exchange the authorization code for a token
	token, err := h.exchangeAuthCode(r)
	if err != nil {
		log.Printf("Failed to exchange authorization code for token: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Use the token to retrieve user information
	userInfo, err := h.fetchUserData(token)
	if err != nil {
		log.Printf("failed to get user info: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if err := h.handleCreateUser(w, userInfo); err != nil {
		return
	}

	// Create JWT token
	jwtToken, err := CreateJWT(userInfo.Email)
	if err != nil {
		log.Printf("Failed to generate JWT: %s\n", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: true,                    // Prevent JavaScript access to cookies
		Secure:   true,                    // Ensure cookies are sent only over HTTPS
		SameSite: http.SameSiteStrictMode, // Prevent CSRF attacks
		Expires:  time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "http://localhost:3000/", http.StatusFound)
}

func (h *Handler) exchangeAuthCode(r *http.Request) (*oauth2.Token, error) {
	rawCode := r.FormValue("code")
	if rawCode == "" {
		log.Println("No auth code provided")
	}

	code, err := url.QueryUnescape(rawCode)
	if err != nil {
		log.Printf("Failed to decode authorization code: %s\n", err)
		
	}

	token, err := config.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Code exchange failed: %s\n", err)
	}

	return token, nil
}



func (h *Handler) fetchUserData(token *oauth2.Token) (*types.UserInfo, error) {
	client := config.GoogleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer response.Body.Close()

	var userInfo types.UserInfo
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to parse user info: %w", err)
	}

	return &userInfo, nil
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Use your existing ValidateToken function
	token, err := ValidateToken(tokenCookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the email from the claims (assuming you stored it as "email" in CreateJWT)
	email, ok := claims["email"].(string) // This is the important part
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusInternalServerError)
		return
	}
	fmt.Println("Extracted email from token:", email)

	// Fetch user data using the email
	user, err := h.userStore.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) handleCreateUser(w http.ResponseWriter, userInfo *types.UserInfo) error {
	var user types.CreateUserPayload
	
	fmt.Println("Email:", userInfo.Email)

	user.FirstName = userInfo.GivenName
	user.LastName = userInfo.FamilyName
	user.Email = userInfo.Email

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	existingUser, err := h.userStore.GetUserByEmail(userInfo.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            // User not found - proceed with creating the user
        } else {
            return fmt.Errorf("failed to check for existing user: %w", err)
        }
    } else {
        log.Printf("User with email %s already exists", existingUser.Email)
        return nil
    }

	err = h.userStore.CreateUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return fmt.Errorf("failed to create user: %w", err)
	}
	utils.WriteJSON(w, http.StatusCreated, user)
	return nil
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