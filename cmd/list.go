package cmd

import (
	"database/sql"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"time"
)

func ListTasks(db *sql.DB) error {
	rows, err := db.Query("SELECT id, task, done, created_at, completed_at FROM tasks ORDER BY id")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Initialize the table writer
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Task", "Status", "Created At", "Completed At"})

	for rows.Next() {
		var id int
		var task string
		var done bool
		var createdAt time.Time
		var completedAt sql.NullTime

		err := rows.Scan(&id, &task, &done, &createdAt, &completedAt)
		if err != nil {
			return err
		}

		status := "❌"
		if done {
			status = "✅"
		}

		completedAtStr := ""
		if completedAt.Valid {
			completedAtStr = completedAt.Time.Format("2006-01-02 15:04")
		}

		t.AppendRow([]interface{}{id, task, status, createdAt.Format("2006-01-02 15:04"), completedAtStr})
	}

	t.Render()
	return nil
}
