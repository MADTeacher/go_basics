package db

import "errors"

func (r *SQLiteRepository) AddTask(task Task, projectID int) (*Task, error) {
	pjTask := &ProjectTask{
		Task:      task,
		ProjectID: projectID,
	}
	tx := r.db.Create(pjTask)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &pjTask.Task, nil
}

func (r *SQLiteRepository) DeleteTask(taskID int) error {
	tx := r.db.Delete(&ProjectTask{Task: Task{ID: taskID}})
	if tx.Error != nil {
		return tx.Error
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return nil
}

func (r *SQLiteRepository) GetAllTasks() (tasks []ProjectTask, err error) {
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

	var pjTasks []ProjectTask
	tx := r.db.Where("project_id", projectID).Find(&pjTasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, ErrNotExists
	}

	for _, it := range pjTasks {
		tasks = append(tasks, it.Task)
	}
	return
}

func (r *SQLiteRepository) TaskDone(taskId int) error {
	if taskId == 0 {
		return errors.New("invalid updated ID")
	}
	pjTask := &ProjectTask{Task: Task{ID: taskId}}
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
