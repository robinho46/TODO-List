package main

import (
	"bufio"
	"fmt"
	"github.com/robinho46/TODO-List/cmd"
	"os"
	"strings"
)

func completedMenu() int {
	fmt.Println("┌────────────────────────────┐")
	fmt.Println("│       Completed Menu       │")
	fmt.Println("│       ─────────────        │")
	fmt.Println("│ 1. Mark task completed     │")
	fmt.Println("│ 2. Delete task             │")
	fmt.Println("│ 3. Undo marked task        │")
	fmt.Println("│ 4. Quit                    │")
	fmt.Println("└────────────────────────────┘")
	fmt.Print("What command do you want to execute: ")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return 0 // or handle error appropriately
	}
	return choice
}

func menu() {
	fmt.Println("Welcome to your TODO list | 📝")
	for {
		fmt.Println("┌────────────────────────────┐")
		fmt.Println("│         TODO List          │")
		fmt.Println("│         ─────────          │")
		fmt.Println("│ 1. Add a task              │")
		fmt.Println("│ 2. Completed a task        │")
		fmt.Println("│ 3. List tasks              │")
		fmt.Println("│ 4. Delete task             │")
		fmt.Println("│ 5. Undo marked task        │")
		fmt.Println("│ 6. Quit                    │")
		fmt.Println("└────────────────────────────┘")
		fmt.Print("What command do you want to execute: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		fmt.Printf("You entered: %d\n", choice)

		switch choice {
		case 1:
			// Add a task
			reader := bufio.NewReader(os.Stdin)
			//fmt.Print("Enter the task description: ")
			fmt.Print("Enter the task description (0 to return): ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			if task == "0" {
				fmt.Println("Returning to the main menu...")
				continue
			}
			err := cmd.AddTask(db, task)
			if err == nil {
				fmt.Println("Added successfully!")
			} else {
				fmt.Println("❌ Add failed:", err)
			}

		case 2:
			// Completed task menu
			subChoice := completedMenu()
			switch subChoice {
			case 1:
				fmt.Print("Enter the id of the task you want to mark as completed (0 to return): ")
				var taskId int
				_, err = fmt.Scan(&taskId)

				if err != nil {
					fmt.Println("❌ Error reading input:", err)
					continue
				}
				if taskId == 0 {
					fmt.Println("Returning to the main menu...")
					continue
				}
				err = cmd.UpdateTask(db, taskId)
				if err != nil {
					fmt.Println("❌ Error when trying to mark as completed:", err)
				} else {
					fmt.Printf("Marked task id %d as completed ✅\n", taskId)
				}
			case 2:
				fmt.Print("Enter the id of the task you want to delete (0 to return): ")
				var taskId int
				_, err = fmt.Scan(&taskId)
				if err != nil {
					fmt.Println("❌ Error reading input:", err)
					continue
				}
				if taskId == 0 {
					fmt.Println("Returning to the main menu...")
					continue
				}
				err = cmd.DeleteTask(db, taskId)
				if err != nil {
					fmt.Println("❌ Error when trying to delete task:", err)
				} else {
					fmt.Printf("Deleted task id %d ✅\n", taskId)
				}
			case 3:
				fmt.Print("Enter the id of the task you want to undo (0 to return): ")
				var taskId int
				_, err = fmt.Scan(&taskId)
				if err != nil {
					fmt.Println("❌ Error reading input:", err)
					continue
				}
				if taskId == 0 {
					fmt.Println("Returning to the main menu...")
					continue
				}
				err = cmd.UndoTask(db, taskId)
				if err != nil {
					fmt.Println("❌ Error when trying to undo task:", err)
				} else {
					fmt.Printf("Unmarked task id %d ✅\n", taskId)
				}
			case 4:
				fmt.Println("Returning to the main menu...")
				continue
			default:
				fmt.Println("❌ Invalid choice in completed menu!")
			}

		case 3:
			// List tasks
			err := cmd.ListTasks(db)
			if err != nil {
				fmt.Println("❌ Error listing tasks:", err)
			}

		case 4:
			fmt.Print("Enter the id of the task you want to delete (0 to return): ")
			var taskId int
			_, err = fmt.Scan(&taskId)
			if err != nil {
				fmt.Println("❌ Error reading input:", err)
				continue
			}
			if taskId == 0 {
				fmt.Println("Returning to the main menu...")
				continue
			}
			err = cmd.DeleteTask(db, taskId)
			if err != nil {
				fmt.Println("❌ Error when trying to delete task:", err)
			} else {
				fmt.Printf("Deleted task id %d ✅\n", taskId)
			}
		case 5:
			fmt.Print("Enter the id of the task you want to undo (0 to return): ")
			var taskId int
			_, err = fmt.Scan(&taskId)
			if err != nil {
				fmt.Println("❌ Error reading input:", err)
				continue
			}
			if taskId == 0 {
				fmt.Println("Returning to the main menu...")
				continue
			}
			err = cmd.UndoTask(db, taskId)
			if err != nil {
				fmt.Println("❌ Error when trying to undo task:", err)
			} else {
				fmt.Printf("Unmarked task id %d ✅\n", taskId)
			}

		case 6:
			fmt.Println("👋 Exiting...")
			return

		default:
			fmt.Println("❌ Invalid choice! Try again.")
		}
	}
}
