package db

import "errors"

// Метод для добавления задачи в базу данных
func (r *SQLiteRepository) AddTask(task Task, projectID int) (*Task, error) {
	pjTask := &ProjectTask{ // создаем связь между задачей и проектом
		Task:      task,
		ProjectID: projectID,
	}
	tx := r.db.Create(pjTask) // создаем задачу
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &pjTask.Task, nil
}

// Метод для удаления задачи из базы данных
func (r *SQLiteRepository) DeleteTask(taskID int) error {
	// удаляем задачу по ее ID
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

// Метод для получения всех задач
func (r *SQLiteRepository) GetAllTasks() (tasks []ProjectTask, err error) {
	tx := r.db.Find(&tasks) // получаем все задачи
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, ErrNotExists
	}

	return
}

// Метод для получения всех задач конкретного проекта
func (r *SQLiteRepository) GetProjectTasks(projectID int) (tasks []Task, err error) {
	if projectID == 0 {
		return nil, errors.New("invalid updated ID")
	}

	var pjTasks []ProjectTask
	// получаем все задачи конкретного проекта
	// и добавляем их в срез pjTasks []ProjectTask
	tx := r.db.Where("project_id", projectID).Find(&pjTasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, ErrNotExists
	}

	for _, it := range pjTasks {
		// Отделяем задачу от проекта
		// и добавляем в срез tasks
		tasks = append(tasks, it.Task)
	}
	return
}

// Метод для обновления задачи
func (r *SQLiteRepository) TaskDone(taskId int) error {
	if taskId == 0 { // Проверка на валидность
		return errors.New("invalid updated ID")
	}
	// Поиск задачи по ее ID. Сначала создаем экземпляр ProjectTask
	// передаем в него Task{ID: taskId}
	pjTask := &ProjectTask{Task: Task{ID: taskId}}
	tx := r.db.Find(&pjTask) // ищем задачу
	if tx.Error != nil {
		return tx.Error
	}
	pjTask.IsDone = true // обновляем поле IsDone
	r.db.Save(&pjTask)   // сохраняем обновленную задачу

	// проверяем обновилась ли задача
	rowsAffected := tx.RowsAffected
	if rowsAffected == 0 {
		return ErrUpdateFailed
	}

	return nil
}
