package service

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type CompletedQuests interface {
	CompleteTask(userID int, questID int) error
}
