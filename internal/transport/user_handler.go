package transport

import (
	"Tasks_Users_Vk_test/internal/model"
	"Tasks_Users_Vk_test/internal/repository"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	Repos *repository.Repositories
}

func NewUserHandler(repos *repository.Repositories) *UserHandler {
	return &UserHandler{
		Repos: repos,
	}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	if user.Name == "" || user.Balance == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "fields are required"})
		return
	}

	err = u.Repos.User.CreateUser(user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
