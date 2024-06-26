package service

import (
	"Tasks_Users_Vk_test/internal/model"
	"Tasks_Users_Vk_test/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type CompletedQuests interface {
	CompleteTask(recordCompleted model.RecordCompleted) error
}

type Services struct {
	CompletedQuests CompletedQuests
}

func NewServices(Repos *repository.Repositories) (*Services, error) {
	return &Services{
		CompletedQuests: NewCompletedQuestsService(Repos),
	}, nil
}
