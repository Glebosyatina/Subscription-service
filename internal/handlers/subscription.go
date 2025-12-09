package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateSub(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "OK")
	w.Write([]byte("OK"))
}

func (h *Handler) GetSub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //мапа с query параметрами
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Sub: ", vars["id"])
}

func (h *Handler) DelSub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Rm Sub: ", vars["id"])
}
