package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"glebosyatina/test_project/internal/domain"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверное тело запроса"))
		return
	}

	u, err := h.services.UserService.AddUser(user.Name, user.Surname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Не удалось создать пользователя")
		return
	}

	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "create user")
	json.NewEncoder(w).Encode(u)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "user:", vars["id"])
}

func (h *Handler) DelUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "delete user:", vars["id"])
}
