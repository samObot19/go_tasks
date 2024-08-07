package service

import (
	"task_manager/models"
	"task_manager/storage"
)

type Services struct{
	Connection Storage.Storage
}

func NewServices(con Storage.Storage) *Services{
	return &Services{Connection: con}
}

func (s *Services) GetTask(id *string) (models.Task, error) {
	task, err := s.Connection.ReadTask(*id)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (s *Services) GetTasks() []models.Task{
	tasks, err := s.Connection.ReadAllTask()
	if err != nil {
		return []models.Task{}
	}
	return tasks
}

func (s *Services) AddTask(newTask *models.Task) error{
	return s.Connection.CreateTask(newTask)
}

func (s *Services) RemoveTask(id *string) error {
	return s.Connection.DeleteTask(*id)
}


func (s *Services) UpdateTask(id *string, updatedTask *models.Task) error {
	return s.Connection.UpdateTask(*id, *updatedTask)
}
