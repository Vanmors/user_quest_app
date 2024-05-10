package transport

import (
	"Tasks_Users_Vk_test/internal/model"
	"Tasks_Users_Vk_test/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

type QuestHandler struct {
	Repos *repository.Repositories
}

func NewQuestHandler(repos *repository.Repositories) *QuestHandler {
	return &QuestHandler{
		Repos: repos,
	}
}

func (q *QuestHandler) CreateQuest(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest
	err := json.NewDecoder(r.Body).Decode(&quest)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	if quest.Name == "" || quest.Cost == 0 || quest.Stages == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "fields are required"})
		return
	}

	err = q.Repos.Quest.CreateQuest(quest)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}
