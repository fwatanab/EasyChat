package database

import (
	"easychat/models"
)

// 新しいユーザーをデータベースに追加
func AddUser(name string) error {
	query := "INSERT INTO list (name, message) VALUES (?, '')"

	_, err := DB.Exec(query, name)
	if err != nil {
		return err
	}
	return nil
}

// 新しいメッセージをデータベースに追加
func AddMessage(userID int64, message string) error {
	query := "UPDATE list SET message = ? WHERE id = ?"

	_, err := DB.Exec(query, message, userID)
	if err != nil {
		return err
	}
	return nil
}

func GetTime(userID int64) (string, error) {
	query := "SELECT timestamp FROM list WHERE id = ?"
	var time string
	err := DB.QueryRow(query, userID).Scan(&time)
	if err != nil {
		return "", err
	}
	return time, nil
}

func GetUserID(name string) (int64, error) {
	query := "SELECT id FROM list WHERE name = ? ORDER BY timestamp DESC LIMIT 1"
	var id int64
	err := DB.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetMessages() ([]models.Message, error) {
	rows, err := DB.Query("SELECT id, name, message, timestamp FROM list ORDER BY timestamp ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var message []models.Message

	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.Name, &msg.InputMessage, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		message = append(message, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return message, nil
}

func DeleteOldMessages(daysAgo int) error {
	query := "DELETE FROM list WHERE timestamp <= DATE_SUB(NOW(), INTERVAL ? DAY)"
	_, err := DB.Exec(query, daysAgo)
	if err != nil {
		return err
	}
	return nil
}
