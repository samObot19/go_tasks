package Services


import (
	"TaskManagementRESTAPI/Models"
	"errors"
	"time"

)

var tasks = [] Models.Task {
	{ID: "1", Title: "Task Manager Project", Description: "Add/View/Delete Tasks", DueDate: time.Now(), Status: "In Progress"},
	{ID: "2", Title: "Books Management Project", Description: "Add/View/Delete Books", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},
}

func GetTask(id *string) (Models.Task, error)  {
	
	for _ , val := range tasks{
		if val.ID == *id{
			return val, nil
		}
	}

	return Models.Task{}, errors.New("not found") 

}

func GetTasks() *[]Models.Task{
	return &tasks
}

func AddTask(newTask *Models.Task) {
	tasks = append(tasks, *newTask)
}

func RemoveTask(id *string) error {

	for index, val := range tasks{
		if val.ID == *id{
			tasks = append(tasks[:index], tasks[index + 1:]...)
			return nil
		}
	}
	return errors.New("not found")
}


func UpdateTask(id *string, updatedTask *Models.Task) error {

	for index, val := range tasks{
		if val.ID == *id{
			if updatedTask.Title != "" {
				tasks[index].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[index].Description = updatedTask.Description
			}

			return nil
		}
	}

	return errors.New("not found")
}