package db

import "errors"

func (r *SQLiteRepository) AddTask(task Task) (*Task, error) {
	tx := r.db.Create(&task)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &task, nil
}

func (r *SQLiteRepository) DeleteTask(taskID int) error {
	tx := r.db.Delete(&Task{ID: taskID})
	if tx.Error != nil {
		return tx.Error
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return nil
}

func (r *SQLiteRepository) GetAllTasks() (tasks []Task, err error) {
	tx := r.db.Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, ErrNotExists
	}

	return
}

func (r *SQLiteRepository) GetProjectTasks(projectID int) (tasks []Task, err error) {
	if projectID == 0 {
		return nil, errors.New("invalid updated ID")
	}

	tx := r.db.Where("project_id", projectID).Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, ErrNotExists
	}

	return
}

func (r *SQLiteRepository) TaskDone(taskId int) error {
	if taskId == 0 {
		return errors.New("invalid updated ID")
	}
	pjTask := &Task{ID: taskId}
	tx := r.db.Find(&pjTask)
	if tx.Error != nil {
		return tx.Error
	}
	pjTask.IsDone = true
	r.db.Save(&pjTask)
	rowsAffected := tx.RowsAffected
	if rowsAffected == 0 {
		return ErrUpdateFailed
	}

	return nil
}
