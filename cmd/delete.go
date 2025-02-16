package cmd

import (
	"database/sql"
	"fmt"
)

func DeleteTask(db *sql.DB, taskId int) error {
	query := "DELETE FROM tasks WHERE id = $1"

	res, err := db.Exec(query, taskId)
	if err != nil {
		return fmt.Errorf("could not delete task as: %v", err)

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
