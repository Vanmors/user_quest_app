package repository

import (
	"Tasks_Users_Vk_test/internal/model"
	"database/sql"
)

type QuestPsql struct {
	conn *sql.DB
}

func NewQuestPsql(db *sql.DB) *QuestPsql {
	return &QuestPsql{
		conn: db,
	}
}

func (q *QuestPsql) GetQuestById(id int) (model.Quest, error) {
	row, err := q.conn.Query("SELECT * FROM quest WHERE id = $1", id)
	if err != nil {
		return model.Quest{}, err
	}

	defer row.Close()

	if !row.Next() {
		return model.Quest{}, sql.ErrNoRows
	}

	quest := model.Quest{}
	err = row.Scan(&quest.Id, &quest.Name, &quest.Cost, &quest.Stages)
	if err != nil {
		return model.Quest{}, err
	}

	return quest, nil
}

func (q *QuestPsql) CreateQuest(quest model.Quest) error {
	_, err := q.conn.Query("INSERT INTO quest (name, cost, stages) VALUES ($1, $2, $3)", quest.Name, quest.Cost, quest.Stages)

	if err != nil {
		return err
	}
	return nil
}

func (q *QuestPsql) GetCost(questID int) (int, error) {
	var cost int
	err := q.conn.QueryRow("SELECT cost FROM quest WHERE id = $1", questID).Scan(&cost)
	if err != nil {
		return 0, err
	}
	return cost, nil
}

func (q *QuestPsql) GetStages(questID int) (int, error) {
	var cost int
	err := q.conn.QueryRow("SELECT stages FROM quest WHERE id = $1", questID).Scan(&cost)
	if err != nil {
		return 0, err
	}
	return cost, nil
}
