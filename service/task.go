package service

import "github.com/leonardo-v-r/golang-to-do-list-api/domain/entity"

type TaskService struct {
	repo entity.TaskRepository
}

func NewTaskService(repo entity.TaskRepository) *TaskService {
	return &TaskService{
		repo,
	}
}

func (s *TaskService) Post(task *entity.Task) error {
	if err := s.repo.Create(task); err != nil {
		return err
	}
	return nil
}
func (s *TaskService) Get() ([]*entity.Task, error)   { return nil, nil }
func (s *TaskService) Update(task *entity.Task) error { return nil }
func (s *TaskService) Delete(ID int) error            { return nil }
