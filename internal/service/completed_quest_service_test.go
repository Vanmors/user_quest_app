package service

//
//import (
//	"Tasks_Users_Vk_test/internal/repository"
//	"Tasks_Users_Vk_test/internal/service"
//	"errors"
//	"testing"
//)
//
//func TestCompletedQuestsService_CompleteTask(t *testing.T) {
//	mockRepos := &repository.MockRepositories{
//		CompletedQuests: &repository.CompletedQuestsPsqlMock{
//			HaveStagesFn: func(userID int, questID int) (int, error) {
//				if questID == 1 {
//					return 0, nil
//				}
//				return 1, nil
//			},
//			AddCompletedTaskFn: func(userID int, questID int) error {
//				return nil
//			},
//			UpdateStagesFn: func(userID int, questID int) error {
//				return nil
//			},
//		},
//		Quest: &repository.QuestPsqlMock{
//			GetCostFn: func(questID int) (int, error) {
//				if questID == 1 {
//					return 10, nil
//				}
//				return 0, errors.New("not found")
//			},
//			GetStagesFn: func(questID int) (int, error) {
//				if questID == 1 {
//					return 2, nil
//				}
//				return 0, errors.New("not found")
//			},
//		},
//		User: &repository.UserPsqlMock{
//			UpdateBalanceFn: func(userID int, amount int) error {
//				return nil
//			},
//		},
//	}
//
//	cs := service.NewCompletedQuestsService(mockRepos)
//
//	t.Run("Quest has been completed", func(t *testing.T) {
//		err := cs.CompleteTask(1, 1)
//		if err == nil || err.Error() != "Quest has been completed" {
//			t.Errorf("Expected 'Quest has been completed' error, got: %v", err)
//		}
//	})
//
//	t.Run("Add completed task", func(t *testing.T) {
//		err := cs.CompleteTask(1, 2)
//		if err != nil {
//			t.Errorf("Expected no error, got: %v", err)
//		}
//	})
//
//	t.Run("Update stages and update balance", func(t *testing.T) {
//		err := cs.CompleteTask(1, 1)
//		if err != nil {
//			t.Errorf("Expected no error, got: %v", err)
//		}
//	})
//}
