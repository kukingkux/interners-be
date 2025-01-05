package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kukingkux/interners-be/service/auth"
	"github.com/kukingkux/interners-be/service/post"
	"github.com/kukingkux/interners-be/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	postStore := post.NewStore(s.db)
	postHandler := post.NewHandler(postStore)
	postHandler.RegisterRoutes(subrouter)

	authStore, err := auth.NewAuthHandler(userStore)
	if err != nil {
		return err
	}
	authStore.RegisterRoutes(subrouter)

	// corsHandler := handlers.CORS(
	// 	handlers.AllowedOrigins([]string{"*"}), // Allow your frontend origin
	// 	handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}), 
	// 	handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	handlers.AllowCredentials(),// This is essential for cookies
	// )
	
	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr,
		handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Allow your frontend origin
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}), 
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),// This is essential for cookies
	)(router))
}
