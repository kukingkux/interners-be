package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kukingkux/interners-be/bin/unused/cart"
	"github.com/kukingkux/interners-be/service/auth"
	"github.com/kukingkux/interners-be/service/order"
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

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(orderStore, postStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	authStore, err := auth.NewAuthHandler(userStore)
	if err != nil {
		return err
	}
	authStore.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
