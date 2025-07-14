package taskservice

import "gorm.io/gorm"

// Основные методы CRUD - Create, Read, Update, Delete

type TaskRepository interface {
	CreateTask(task Task) error
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(id string) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepository) UpdateTask(task Task) error {
	return r.db.Save(&task).Error
}

func (r *taskRepository) DeleteTask(id string) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
