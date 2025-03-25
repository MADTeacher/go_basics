package db

import "errors"

func (s *SQLiteRepository) AddTask(task Task, projectID int) (*Task, error) {
	res, err := s.db.Exec("INSERT INTO tasks(name, description, priority, is_done, project_id) values(?,?,?,?,?)",
		task.Name, task.Description, task.Priority, task.IsDone, projectID)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	task.ID = int(id)

	return &task, nil
}

func (s *SQLiteRepository) DeleteTask(taskID int) error {
	res, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", taskID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	// RowsAffected returns the number of rows affected by an update, insert, or delete.
	// Not every database or database driver may support this.
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}

func (s *SQLiteRepository) GetAllTasks() (tasks []ProjectTask, err error) {
	rows, err := s.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task ProjectTask
		if err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Priority,
			&task.IsDone, &task.ProjectID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return
}

func (s *SQLiteRepository) GetProjectTasks(projectID int) (tasks []Task, err error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE project_id = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		var progID int
		if err := rows.Scan(&task.ID, &task.Name, &task.Description,
			&task.Priority, &task.IsDone, &progID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return
}

func (s *SQLiteRepository) TaskDone(taskId int) error {
	if taskId == 0 {
		return errors.New("invalid updated ID")
	}
	res, err := s.db.Exec("UPDATE tasks SET is_done = ? WHERE id = ?", 1, taskId)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrUpdateFailed
	}

	return nil
}
