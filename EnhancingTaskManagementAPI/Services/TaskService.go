package Service

import (
	"EnhancingTaskManagementAPI/Models"
	"EnhancingTaskManagementAPI/Storage"
)

type Services struct{
	Connection Storage.Storage
}

func NewServices(con Storage.Storage) *Services{
	return &Services{Connection: con}
}

func (s *Services) GetTask(id *string) (Models.Task, error) {
	task, err := s.Connection.Read(*id)
    if err != nil {
        return Models.Task{}, err
    }
    return task.(Models.Task), nil

}

func (s *Services) GetTasks() []Models.Task{
	tasks, err := s.Connection.ReadAll()
    if err != nil {
        return []Models.Task{}
    }
    return tasks.([]Models.Task)

}

func (s *Services) AddTask(newTask *Models.Task) {
	s.Connection.Create(newTask)
}

func (s *Services) RemoveTask(id *string) error {
	return s.Connection.Delete(*id)
}


func (s *Services) UpdateTask(id *string, updatedTask *Models.Task) error {
	return s.Connection.Update(*id, updatedTask)
}
