package service

import (
	"Tasks_Users_Vk_test/internal/domain"
	"Tasks_Users_Vk_test/internal/repository"
	"errors"
)

type CompletedQuestsService struct {
	repos *repository.Repositories
}

func NewCompletedQuestsService(Repos *repository.Repositories) *CompletedQuestsService {
	return &CompletedQuestsService{
		repos: Repos,
	}
}

func (cs *CompletedQuestsService) CompleteTask(recordCompleted domain.RecordCompleted) error {
	userID := recordCompleted.UserID
	questID := recordCompleted.QuestID

	haveStages, err := cs.repos.CompletedQuests.HaveStages(userID, questID)
	if err != nil {
		return err
	}
	questCost, err := cs.repos.Quest.GetCost(questID)
	if err != nil {
		return err
	}
	needStages, err := cs.repos.Quest.GetStages(questID)
	if err != nil {
		return err
	}

	if haveStages == needStages {
		return errors.New("Quest has been completed")
	}

	if haveStages+1 == needStages {
		cs.repos.User.UpdateBalance(userID, questCost)
	}

	if haveStages == 0 {
		cs.repos.CompletedQuests.AddCompletedTask(userID, questID)
	} else {
		cs.repos.CompletedQuests.UpdateStages(userID, questID)
	}
	return nil
}
