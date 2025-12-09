package handlers

import (
	"log/slog"

	"github.com/gorilla/mux"

	"glebosyatina/test_project/internal/service"
)

// слой обработки запросов, в качестве зависимостей включает слой service и логгер
type Handler struct {
	services *service.Services
	logger   *slog.Logger
}

func NewHandler(servs *service.Services, logg *slog.Logger) *Handler {
	return &Handler{
		services: servs,
		logger:   logg,
	}
}

func (h *Handler) InitRoutes() *mux.Router {

	r := mux.NewRouter()

	subs := r.PathPrefix("/sub").Subrouter()
	subs.HandleFunc("/add", h.CreateSub).Methods("POST")
	subs.HandleFunc("/{id}", h.GetSub).Methods("GET")
	subs.HandleFunc("/rm/{id}", h.DelSub).Methods("DELETE")

	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/add", h.CreateUser).Methods("POST")
	users.HandleFunc("/{id}", h.GetUser).Methods("GET")
	users.HandleFunc("/rm/{id}", h.DelUser).Methods("GET")

	return r
}
