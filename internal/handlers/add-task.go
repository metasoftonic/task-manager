package handlers

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/metasoftonic/task-manager/internal/helpers"
	"github.com/metasoftonic/task-manager/internal/models"
)

func AddCommand(scanner *bufio.Scanner, taskList *[]models.Task) (task models.Task, err error) {
	fmt.Print("What is your name?\n")
	scanner.Scan()
	name := strings.ToLower(scanner.Text())

	fmt.Print("What is the task title?\n")
	scanner.Scan()
	title := strings.ToLower(scanner.Text())

	fmt.Print("What is the task summary or description?\n")
	scanner.Scan()
	description := strings.ToLower(scanner.Text())

	fmt.Print("What is the current status of the task?: Todo, In progress, Done\n")
	scanner.Scan()
	Status := strings.ToLower(scanner.Text())

	fmt.Print("What is the state date for this task?\n")
	scanner.Scan()
	stateDate, err := helpers.ParseDate(scanner.Text())
	if err != nil {
		return models.Task{}, errors.New("error passing start date string to date. use formate yyyy-mm-dd")
	}
	fmt.Print("What is the end date for this task?\n")
	scanner.Scan()
	endDate, err := helpers.ParseDate(scanner.Text())
	if err != nil {
		return models.Task{}, errors.New("error passing end date string to date. use formate yyyy-mm-dd")
	}

	if stateDate.After(endDate) {
		return models.Task{}, errors.New("end date must be greater than start date")
	}
	t := models.Task{
		Id:          uuid.New().String(),
		Username:    name,
		Title:       title,
		Description: description,
		Status:      Status,
		StartDate:   stateDate,
		EndDate:     endDate,
		CreatedDate: time.Now(),
	}
	*taskList = append(*taskList, t)
	printed := fmt.Sprintf("Task count: %d, Recent task ID: %s", len(*taskList), t.Id)
	fmt.Println(printed)
	return t, nil
}
