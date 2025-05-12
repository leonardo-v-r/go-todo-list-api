package entity

import (
	"time"
)

type TaskStatus int

const (
	Pending = iota
	Completed
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
}

type TaskRepository interface {
	Create(task *Task) error
	Read() ([]*Task, error)
	Update(task *Task) error
	Delete(ID int) error
}

type TaskService interface {
	Post(task *Task) error
	Get() ([]*Task, error)
	Update(task *Task) error
	Delete(ID int) error
}
