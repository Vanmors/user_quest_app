package domain

import "time"

type CompletedQuests struct {
	Id      int
	UserId  int
	QuestId int
	Stages  int
	Date    time.Time
}
