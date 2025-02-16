package cmd

import (
	"database/sql"
	"errors"
)

func AddTask(db *sql.DB, task string) error {
	if task == "" {
		return errors.New("You gave me an empty task!")
	}
	query := "INSERT INTO tasks (task) VALUES ($1)"

	_, err := db.Exec(query, task)
	if err != nil {
		return err
	}
	return nil
}
