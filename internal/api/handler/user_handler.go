package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/Quszlet/CRUD-operations/internal/models"
	"github.com/Quszlet/CRUD-operations/pkg/JSON"
	"github.com/gorilla/mux"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := JSON.Parse(r, user)
	if err != nil {
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Failed parse JSON")
		return
	}

	err = user.Validate()
	if err != nil {
		JSON.ErrorResponse(w, http.StatusBadRequest, err.Error(), "Invalid JSON")
		return
	}

	id, err := h.services.Create(user)
	if err != nil {
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed created user")
		return
	}

	message := fmt.Sprintf("User created with id %d", id)

	JSON.Response(w, http.StatusCreated, message)
}

// Подумать как обновлять поля, которые указаны в JSON
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	//id, _ := strconv.Atoi(mux.Vars(r)["id"])
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.services.Get(id)
	if err != nil {
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed get user")
		return
	}

	JSON.Response(w, http.StatusCreated, user)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.services.GetAll()
	if err != nil {
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed get users")
		return
	}

	JSON.Response(w, http.StatusCreated, users)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := h.services.Delete(id)
	if err != nil {
		JSON.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "Failed delete user")
		return
	}

	message := fmt.Sprintf("User delete with id %d", id)
	
	JSON.Response(w, http.StatusCreated, message)
}
