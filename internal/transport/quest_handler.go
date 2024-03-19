package transport

import (
	"Tasks_Users_Vk_test/internal/domain"
	"Tasks_Users_Vk_test/internal/repository"
	"Tasks_Users_Vk_test/pkg/util"
	"encoding/json"
	"github.com/gorilla/mux"
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

func (q *QuestHandler) GetQuest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questId := util.MustAtoi(vars["id"])
	//quest, err := q.Repos.Quest.GetQuestById(questId)
	cost, err := q.Repos.Quest.GetCost(questId)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(cost)
}

func (q *QuestHandler) CreateQuest(w http.ResponseWriter, r *http.Request) {
	var quest domain.Quest
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
