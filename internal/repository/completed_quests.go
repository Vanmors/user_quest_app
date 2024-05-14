package repository

import (
	"Tasks_Users_Vk_test/internal/model"
	"database/sql"
)

type CompletedQuestsPsql struct {
	conn *sql.DB
}

func NewCompletedQuestsPsql(db *sql.DB) *CompletedQuestsPsql {
	return &CompletedQuestsPsql{
		conn: db,
	}
}

func (c *CompletedQuestsPsql) HaveStages(userID int, questID int) (int, error) {
	var completed int
	err := c.conn.QueryRow("SELECT stages FROM completed_quests WHERE user_id = $1 AND quest_id = $2", userID, questID).Scan(&completed)
	if err != nil {
		// Если не строка не найдена, то stage равен 0
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return completed, nil
}

func (c *CompletedQuestsPsql) UpdateStages(userID int, questID int) error {
	err := c.conn.QueryRow("UPDATE completed_quests SET stages = stages + 1 WHERE user_id = $1 AND quest_id = $2", userID, questID)
	if err != nil {
		return sql.ErrNoRows
	}
	return err.Err()
}

func (c *CompletedQuestsPsql) AddCompletedTask(userID int, questID int) error {
	_, err := c.conn.Query("INSERT INTO completed_quests (user_id, quest_id, stages) VALUES ($1, $2, 1)", userID, questID)

	if err != nil {
		return err
	}
	return err
}

func (c *CompletedQuestsPsql) GetCompletedQuestsByUserId(userID int) ([]model.HistoryQuests, error) {
	var historyQuests []model.HistoryQuests

	rows, err := c.conn.Query("SELECT c.user_id, c.quest_id, c.Stages, c.completion_date, q.Name AS QuestName, q.Cost FROM completed_quests c JOIN quest q ON c.quest_id = q.id WHERE c.user_id = $1 ORDER BY c.completion_date DESC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Итерируем по результатам запроса и добавляем выполненные задания в список
	for rows.Next() {
		var historyQuest model.HistoryQuests
		if err := rows.Scan(&historyQuest.UserID, &historyQuest.QuestID, &historyQuest.Stages,
			&historyQuest.CompletionDate, &historyQuest.QuestName, &historyQuest.Cost); err != nil {
			return nil, err
		}
		historyQuests = append(historyQuests, historyQuest)
	}

	return historyQuests, nil
}
