package cmd

import (
	"database/sql"
	"fmt"
)

func UpdateTask(db *sql.DB, taskId int) error {
	query := "UPDATE tasks SET completed_at = NOW(), done = TRUE WHERE id = $1 AND done = FALSE"

	res, err := db.Exec(query, taskId)
	if err != nil {
		return fmt.Errorf("could not mark task as completed: %v", err)
	}

	// Check how many rows were affected
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not retrieve affected rows: %v", err)
	}

	if count == 0 {
		return fmt.Errorf("no task found with id %d", taskId)
	}

	return nil

}

func UndoTask(db *sql.DB, taskId int) error {
	query := "UPDATE tasks SET completed_at = NULL, done = FALSE WHERE id = $1 AND done = TRUE"

	res, err := db.Exec(query, taskId)
	if err != nil {
		return fmt.Errorf("could not undo task completion: %v", err)
	}

	// Check how many rows were affected
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not retrieve affected rows: %v", err)
	}

	if count == 0 {
		return fmt.Errorf("no completed task found with id %d", taskId)
	}

	return nil

}
