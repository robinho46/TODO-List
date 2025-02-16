package main

import (
	"bufio"
	"fmt"
	"github.com/robinho46/TODO-List/cmd"
	"os"
	"strings"
)

func main() {
	connectDB()
	defer db.Close()

	//fmt.Println("📝 TODO List CLI is running!")

	for {
		fmt.Println("Welcome to your TODO list | 📝")
		fmt.Println("------------------------------")
		fmt.Println("| 1. Add a task              |")
		fmt.Println("| 2. Completed a task        |")
		fmt.Println("| 3. List tasks              |")
		fmt.Println("| 4. Quit                    |")
		fmt.Println("------------------------------")
		fmt.Println("What do you want to do: ")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
		}
		fmt.Printf("You entered: %d\n", choice)

		switch choice {
		case 1:
			reader := bufio.NewReader(os.Stdin)
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			err := cmd.AddTask(db, task)
			if err == nil {
				fmt.Println("Added successfully!")
			} else {
				fmt.Println("Add failed?")
			}
		case 2:
			break
		case 3:
			err := cmd.ListTasks(db)
			if err != nil {
				fmt.Println("Error listing tasks:", err)
			} else {
				continue
			}
		case 4:
			fmt.Println("👋 Exiting...")
			return

		default:
			fmt.Println("❌ Invalid choice! Try again.")

		}
	}
}
