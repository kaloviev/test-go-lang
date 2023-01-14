package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alsolovyev/dummy-api/internal/entity"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u entity.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()

	if err := v.Struct(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UserService.Create(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
