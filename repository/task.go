package repository

import "database/sql"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type TaskRepository interface {
	Create(task *Task) (int, error)
	Read(id int) (*Task, error)
	Update(task *Task) error
	Delete(id int) error
}

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepositoryImpl{db: db}
}

func (t taskRepositoryImpl) Create(task *Task) (int, error) {
	stmt := `INSERT INTO tasks (title) VALUES (?) RETURNING id`
	err := t.db.QueryRow(stmt, task.Title).Scan(&task.ID)
	return task.ID, err
}

func (t taskRepositoryImpl) Read(id int) (*Task, error) {
	stmt := `SELECT id, title FROM tasks WHERE id = ?`
	task := &Task{}
	err := t.db.QueryRow(stmt, id).Scan(&task.ID, &task.Title)
	return task, err
}

func (t taskRepositoryImpl) Update(task *Task) error {
	stmt := `UPDATE tasks SET title = ? WHERE id = ?`
	result, err := t.db.Exec(stmt, task.Title, task.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (t taskRepositoryImpl) Delete(id int) error {
	stmt := `DELETE FROM tasks WHERE id = ?`
	result, err := t.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
