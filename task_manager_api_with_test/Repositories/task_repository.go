package repository


import (
	"github.com/task_manager/Domain"
)

type TaskRepository interface{
	CreateTask(data *domain.Task) error
	ReadTask(id string) (domain.Task, error)
	ReadAllTask() ([]domain.Task, error)
	UpdateTask(id string, data domain.Task) error
	DeleteTask(id string) error
}