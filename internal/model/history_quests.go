package model

import "time"

type HistoryQuests struct {
	UserID         int
	QuestID        int
	Stages         int
	CompletionDate time.Time
	QuestName      string
	Cost           int
}
