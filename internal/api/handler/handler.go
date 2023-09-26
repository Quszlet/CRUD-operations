package handler

import (
	"github.com/gorilla/mux"
	"github.com/Quszlet/CRUD-operations/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/create", h.Create).Methods("POST")
	users.HandleFunc("/update/{id:[0-9]+}", h.Update).Methods("UPDATE")
	users.HandleFunc("/{id:[0-9]+}", h.Get).Methods("GET")
	users.HandleFunc("", h.GetAll).Methods("GET")
	users.HandleFunc("/delete/{id:[0-9]+}", h.Delete).Methods("DELETE")
	return r
}