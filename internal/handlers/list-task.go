package handlers

import (
	"os"

	"github.com/metasoftonic/task-manager/internal/models"
	"github.com/olekukonko/tablewriter"
)

func PrintTaskListTable(tasks *[]models.Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Username", "Title", "Description", "Status", "Start Date", "End Date", "Created Date"})

	for _, task := range *tasks {
		table.Append([]string{
			task.Id,
			task.Username,
			task.Title,
			task.Description,
			task.Status,
			task.StartDate.Format("2006-01-02"),
			task.EndDate.Format("2006-01-02"),
			task.CreatedDate.Format("2006-01-02"),
		})
	}

	table.Render()
}
