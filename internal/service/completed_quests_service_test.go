package service

import (
	"Tasks_Users_Vk_test/internal/model"
	"Tasks_Users_Vk_test/internal/repository"
	mock_repository "Tasks_Users_Vk_test/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompletedQuestsService_CompleteTask_Error(t *testing.T) {
	// Создание контроллера gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создание мок-объектов для User, Quest и CompletedQuests
	mockUser := mock_repository.NewMockUser(ctrl)
	mockQuest := mock_repository.NewMockQuest(ctrl)
	mockCompletedQuests := mock_repository.NewMockCompletedQuests(ctrl)

	service := Services{CompletedQuests: NewCompletedQuestsService(&repository.Repositories{
		User:            mockUser,
		Quest:           mockQuest,
		CompletedQuests: mockCompletedQuests,
	})}

	// Задание ожиданий для вызовов методов мок-объектов
	mockCompletedQuests.EXPECT().HaveStages(gomock.Any(), gomock.Any()).Return(0, nil)
	mockQuest.EXPECT().GetCost(gomock.Any()).Return(100, nil)
	mockQuest.EXPECT().GetStages(gomock.Any()).Return(3, nil)
	mockUser.EXPECT().UpdateBalance(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockCompletedQuests.EXPECT().AddCompletedTask(gomock.Any(), gomock.Any()).Return(nil)

	// Выполнение тестируемой функции
	err := service.CompletedQuests.CompleteTask(model.RecordCompleted{
		UserID:  1,
		QuestID: 1,
	})

	// Проверка ожидаемого результата
	assert.NoError(t, err)
}

func TestCompletedQuestsService_CompleteTask_OK(t *testing.T) {
	// Создание контроллера gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создание мок-объектов для User, Quest и CompletedQuests
	mockUser := mock_repository.NewMockUser(ctrl)
	mockQuest := mock_repository.NewMockQuest(ctrl)
	mockCompletedQuests := mock_repository.NewMockCompletedQuests(ctrl)

	service := Services{CompletedQuests: NewCompletedQuestsService(&repository.Repositories{
		User:            mockUser,
		Quest:           mockQuest,
		CompletedQuests: mockCompletedQuests,
	})}

	// Задание ожиданий для вызовов методов мок-объектов
	mockCompletedQuests.EXPECT().HaveStages(gomock.Any(), gomock.Any()).Return(3, nil)
	mockQuest.EXPECT().GetCost(gomock.Any()).Return(100, nil)
	mockQuest.EXPECT().GetStages(gomock.Any()).Return(3, nil)
	mockUser.EXPECT().UpdateBalance(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockCompletedQuests.EXPECT().AddCompletedTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	// Выполнение тестируемой функции
	err := service.CompletedQuests.CompleteTask(model.RecordCompleted{
		UserID:  1,
		QuestID: 1,
	})

	// Проверка ожидаемого результата
	assert.EqualError(t, err, "Quest has been completed")
}
