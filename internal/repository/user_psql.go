package repository

import (
	"Tasks_Users_Vk_test/internal/model"
	"database/sql"
)

type UserPsql struct {
	conn *sql.DB
}

func NewUserPsql(db *sql.DB) *UserPsql {
	return &UserPsql{
		conn: db,
	}
}

func (u *UserPsql) GetUserById(id int) (model.User, error) {
	rows, err := u.conn.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return model.User{}, sql.ErrNoRows
	}

	user := model.User{}
	err = rows.Scan(&user.Id, &user.Name, &user.Balance)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UserPsql) CreateUser(user model.User) error {
	_, err := u.conn.Query("INSERT INTO users (name, balance) VALUES ($1, $2)", user.Name, user.Balance)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserPsql) UpdateBalance(userID int, questCost int) error {
	err := u.conn.QueryRow("UPDATE users SET balance = balance + $1 WHERE id = $2", questCost, userID)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (u *UserPsql) GetBalance(userID int) (int, error) {
	var balance int
	err := u.conn.QueryRow("SELECT balance FROM users WHERE id = $1", userID).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
