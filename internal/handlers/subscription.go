package handlers

import (
	"encoding/json"
	"fmt"
	"glebosyatina/test_project/internal/domain"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateSub(w http.ResponseWriter, r *http.Request) {

	var sub domain.Sub
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неверное тело запроса")
		return
	}

	s, err := h.services.SubService.AddSub(sub.UserId, sub.NameService, sub.Price, sub.Start, sub.End)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Не удалось добавить подписку")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}

func (h *Handler) GetSub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //мапа с query параметрами
	subId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неверный id в строке запроса")
		return
	}

	sub, err := h.services.SubService.GetSubscription(uint64(subId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Не удалось получить подписку")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sub)
}

func (h *Handler) GetSubs(w http.ResponseWriter, r *http.Request) {
	subs, err := h.services.SubService.GetSubs()
	if err != nil {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Не удалось получить список подписок")
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subs)
}

func (h *Handler) DelSub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неверный id в строке запроса")
		return
	}

	err = h.services.SubService.DeleteSubByID(uint64(subId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Не удалось удалить подписку")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Удалена подписка: %d", subId)
}

func (h *Handler) UpdateSub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неверно передан параметр")
		return
	}

	var sub domain.Sub
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверное тело запроса"))
		return
	}

	subscription, err := h.services.SubService.UpdateSub(uint64(id), sub.UserId, sub.NameService, sub.Price, sub.Start, sub.End)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Не удалось обновить подписку")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subscription)
}
