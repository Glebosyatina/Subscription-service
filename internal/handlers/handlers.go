package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"

	"glebosyatina/test_project/internal/service"
	"glebosyatina/test_project/internal/handlers/middleware"
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

func (h *Handler) InitRoutes() http.Handler {

	r := mux.NewRouter()

	subs := r.PathPrefix("/sub").Subrouter()
	subs.HandleFunc("/add", h.CreateSub).Methods("POST")
	subs.HandleFunc("/{id}", h.GetSub).Methods("GET")
	subs.HandleFunc("/", h.GetSubs).Methods("GET")
	subs.HandleFunc("/rm/{id}", h.DelSub).Methods("DELETE")
	subs.HandleFunc("/update/{id}", h.UpdateSub).Methods("PUT")

	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/add", h.CreateUser).Methods("POST")
	users.HandleFunc("/{id}", h.GetUser).Methods("GET")
	users.HandleFunc("/", h.GetUsers).Methods("GET")
	users.HandleFunc("/rm/{id}", h.DelUser).Methods("DELETE")
	users.HandleFunc("/update/{id}", h.UpdateUser).Methods("PUT")

	//global middleware
	handler := middleware.Logging(r)
	
	return handler
}


