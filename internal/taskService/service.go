package taskservice

import (
	"errors"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(req TaskRequest) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, req TaskRequest) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) validateTaskRequest(req TaskRequest) error {
	if req.Task == "" {
		return errors.New("task cannot be empty")
	}
	return nil
}

func (s *taskService) CreateTask(req TaskRequest) (Task, error) {
	if err := s.validateTaskRequest(req); err != nil {
		return Task{}, err
	}

	task := Task{
		ID:     uuid.New().String(),
		Task:   req.Task,
		IsDone: req.IsDone,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
	if id == "" {
		return Task{}, errors.New("task ID is required")
	}
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id string, req TaskRequest) (Task, error) {
	if id == "" {
		return Task{}, errors.New("task ID is required")
	}

	if err := s.validateTaskRequest(req); err != nil {
		return Task{}, err
	}

	task := Task{
		ID:     id,
		Task:   req.Task,
		IsDone: req.IsDone,
	}

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	if id == "" {
		return errors.New("task ID is required")
	}
	return s.repo.DeleteTask(id)
}
