package repository

import (
	"database/sql"
	"time"

	"github.com/leonardo-v-r/golang-to-do-list-api/domain/entity"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db,
	}
}

func (r *TaskRepository) Create(task *entity.Task) error {
	stmt, err := r.db.Prepare(`insert into task_tb (task_title, task_description, task_status, created_at) 
		values (?,?,?,?) 
		returning task_id, task_title, task_description, task_status, created_at
	`)
	if err != nil {
		return err
	}
	row := stmt.QueryRow(task.Title, task.Description, task.Status, time.Now().Format("2006-01-02"))
	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt); err != nil {
		return err
	}
	return nil
}
func (r *TaskRepository) Read() ([]*entity.Task, error)  { return nil, nil }
func (r *TaskRepository) Update(task *entity.Task) error { return nil }
func (r *TaskRepository) Delete(ID int) error            { return nil }
