package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/metasoftonic/task-manager/internal/handlers"
	"github.com/metasoftonic/task-manager/internal/models"
)

func callCommand(taskList *[]models.Task, scanner *bufio.Scanner) error {

	stop := false
	for {
		if stop {
			break
		}

		fmt.Print("What do you wish to do?\n")

		scanner.Scan()

		command := strings.ToLower(scanner.Text())

		switch command {
		case "add":
			_, err := handlers.AddCommand(scanner, taskList)
			if err != nil {
				return err
			}
		case "list":
			handlers.PrintTaskListTable(taskList)
		case "update":
			err := handlers.UpdateCommand(scanner, taskList)
			if err != nil {
				return err

			}
		case "exit":
			fmt.Println("Exiting application...")
			time.Sleep(2 * time.Second)
			stop = true
		default:
			return errors.New("unrecognized command")
		}
	}
	return nil

}
func main() {
	// LIST COMMANDS TO USER
	var taskList []models.Task
	fmt.Println("The following commands exist for managing task.")
	fmt.Println("Add => Adds a new task.\n List => list all task created.\n Update => updates a task.\n Delete => deletes a task using the task ID")
	//PROMPT THEM TO CHOSE A COMMAND
	scanner := bufio.NewScanner(os.Stdin)
	var err error
	err = callCommand(&taskList, scanner)
	for err != nil {
		fmt.Println("errors:", err)
		err = callCommand(&taskList, scanner)
	}
	//SCAN THE INPUT
}
