package transport

import (
	"Tasks_Users_Vk_test/internal/domain"
	"Tasks_Users_Vk_test/internal/repository"
	"Tasks_Users_Vk_test/internal/service"
	"Tasks_Users_Vk_test/pkg/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CompletedQuestsHandler struct {
	Repos   *repository.Repositories
	Service *service.Services
}

func NewCompletedQuestsHandler(repos *repository.Repositories, serv *service.Services) *CompletedQuestsHandler {
	return &CompletedQuestsHandler{
		Repos:   repos,
		Service: serv,
	}
}

func (ch *CompletedQuestsHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {

	var recordCompleted domain.RecordCompleted
	err := json.NewDecoder(r.Body).Decode(&recordCompleted)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	if recordCompleted.UserID == 0 || recordCompleted.QuestID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "fields are required"})
		return
	}

	err = ch.Service.CompletedQuests.CompleteTask(recordCompleted)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

func (ch *CompletedQuestsHandler) GetCompletedQuestsAndBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := util.MustAtoi(vars["id"])
	historyQuest, err := ch.Repos.CompletedQuests.GetCompletedQuestsByUserId(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	balance, err := ch.Repos.User.GetBalance(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(balance)
	json.NewEncoder(w).Encode(historyQuest)
}
