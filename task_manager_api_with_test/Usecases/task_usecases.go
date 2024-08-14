package usecase

import (
	"github.com/task_manager/Repositories"
	"github.com/task_manager/Domain"
)

type TaskUsecase struct{
	db repository.TaskRepository
}

func NewTaskUsecase(con repository.TaskRepository) *TaskUsecase{
	return &TaskUsecase{db: con}
}

func (s *TaskUsecase) GetTask(id *string) (domain.Task, error) {
	task, err := s.db.ReadTask(*id)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (s *TaskUsecase) GetTasks() []domain.Task{
	tasks, err := s.db.ReadAllTask()
	if err != nil {
		return []domain.Task{}
	}
	return tasks
}

func (s *TaskUsecase) AddTask(newTask *domain.Task) error{
	return s.db.CreateTask(newTask)
}

func (s *TaskUsecase) RemoveTask(id *string) error {
	return s.db.DeleteTask(*id)
}


func (s *TaskUsecase) UpdateTask(id *string, updatedTask *domain.Task) error {
	return s.db.UpdateTask(*id, *updatedTask)
}
