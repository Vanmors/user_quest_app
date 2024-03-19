package repository

import (
	"Tasks_Users_Vk_test/internal/domain"
	"Tasks_Users_Vk_test/pkg/store"
	"fmt"
	"golang.org/x/text/encoding/charmap"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type User interface {
	GetUserById(id int) (domain.User, error)
	CreateUser(user domain.User) error
	UpdateBalance(userID int, questCost int) error
	GetBalance(userID int) (int, error)
}

type Quest interface {
	GetQuestById(id int) (domain.Quest, error)
	CreateQuest(quest domain.Quest) error
	GetCost(questID int) (int, error)
	GetStages(questID int) (int, error)
}

type CompletedQuests interface {
	HaveStages(userID int, questsID int) (int, error)
	AddCompletedTask(userID int, questID int) error
	UpdateStages(userID int, questID int) error
	GetCompletedQuestsByUserId(userID int) ([]domain.HistoryQuests, error)
}

type Repositories struct {
	User            User
	Quest           Quest
	CompletedQuests CompletedQuests
}

func NewRepositories(dbname, username, password, host, port string) (*Repositories, error) {
	db, err := store.NewClient(dbname, username, password, host, port)
	if err != nil {
		return nil, wrapErrorFromDB(err)
	}
	return &Repositories{
		User:            NewUserPsql(db),
		Quest:           NewQuestPsql(db),
		CompletedQuests: NewCompletedQuestsPsql(db),
	}, nil
}

func wrapErrorFromDB(err error) error {
	if err == nil {
		return err
	}
	utf8Text, _ := charmap.Windows1251.NewDecoder().String(err.Error())
	return fmt.Errorf(utf8Text)
}
