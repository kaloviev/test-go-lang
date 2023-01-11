package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alsolovyev/dummy-api/internal/entity"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u entity.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n", u)
}
