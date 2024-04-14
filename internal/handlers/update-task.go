package handlers

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/metasoftonic/task-manager/internal/helpers"
	"github.com/metasoftonic/task-manager/internal/models"
)

func UpdateCommand(scanner *bufio.Scanner, taskList *[]models.Task) error {
	//ask for task id
	fmt.Print("What is the task Id?\n")
	scanner.Scan()
	taskId := strings.ToLower(scanner.Text())

	var task *models.Task
	for i, t := range *taskList {
		if t.Id == taskId {
			task = &(*taskList)[i]
			break
		}
	}

	if task.Id == "" {
		return errors.New("no task record found for task id " + taskId)
	}
	fmt.Print("What task properties do you wish to update(comma seperated property name)?\n")
	scanner.Scan()
	propertyString := strings.ToLower(scanner.Text())

	fmt.Print("What is the value of the properties (comma seperated to match the property name)?\n")
	scanner.Scan()
	propertyValString := strings.ToLower(scanner.Text())

	propertyArr := strings.Split(propertyString, ",")
	propertyVal := strings.Split(propertyValString, ",")

	if len(propertyArr) != len(propertyVal) {
		return errors.New("property name length should match property value length")
	}

	for i := 0; i < len(propertyArr); i++ {
		err := fillProperty(task, propertyArr[i], propertyVal[i])
		if err != nil {
			return err
		}
	}

	if task.StartDate.After(task.EndDate) {
		return errors.New("end date must be greater than start date")
	}

	return nil
}

func fillProperty(updateTask *models.Task, propertyName string, propertyValue string) error {

	switch strings.ToLower(propertyName) {
	case "title":
		updateTask.Title = propertyValue
	case "description":
		updateTask.Description = propertyValue
	case "status":
		updateTask.Status = propertyValue
	case "startdate":
		startDate, err := helpers.ParseDate(propertyValue)
		if err != nil {
			return err
		}
		updateTask.StartDate = startDate
	case "enddate":
		endDate, err := helpers.ParseDate(propertyValue)
		if err != nil {
			return err
		}
		updateTask.EndDate = endDate

	}
	return nil
}
