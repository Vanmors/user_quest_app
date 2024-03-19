package transport

import (
	"Tasks_Users_Vk_test/internal/repository"
	"Tasks_Users_Vk_test/internal/service"
	"Tasks_Users_Vk_test/pkg/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CompletedQuestsHandler struct {
	Repos   *repository.Repositories
	Service *service.CompletedQuestsService
}

func NewCompletedQuestsHandler(repos *repository.Repositories, serv *service.CompletedQuestsService) *CompletedQuestsHandler {
	return &CompletedQuestsHandler{
		Repos:   repos,
		Service: serv,
	}
}

func (ch *CompletedQuestsHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	type CompletedTask struct {
		UserID  int `json:"userID"`
		QuestID int `json:"questID"`
	}
	var completedTask CompletedTask
	json.NewDecoder(r.Body).Decode(&completedTask)

	err := ch.Service.CompleteTask(completedTask.UserID, completedTask.QuestID)
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
