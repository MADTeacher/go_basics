package db

import "errors"

// Метод для добавления задачи в базу данных
func (s *SQLiteRepository) AddTask(task Task, projectID int) (*Task, error) {
	// Запрос на добавление задачи  проекту с id == projectID
	res, err := s.db.Exec(
		"INSERT INTO tasks(name, description, priority,"+
			" is_done, project_id) values(?,?,?,?,?)",
		task.Name, task.Description, task.Priority,
		task.IsDone, projectID,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId() // Получаем ID
	if err != nil {
		return nil, err
	}
	task.ID = int(id) // Устанавливаем ID задаче

	return &task, nil
}

// Метод для удаления задачи из базы данных
func (s *SQLiteRepository) DeleteTask(taskID int) error {
	// Запрос на удаление задачи из таблицы tasks
	res, err := s.db.Exec(
		"DELETE FROM tasks WHERE id = ?",
		taskID,
	)
	if err != nil {
		return err
	}

	// Проверяем, была ли удалена задача
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// Если не была удалена ни одна задача
	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}

// Метод для получения всех задач
func (s *SQLiteRepository) GetAllTasks() (tasks []ProjectTask, err error) {
	// Запрос на получение всех задач
	rows, err := s.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task ProjectTask
		// считываем данные из каждой строки, в соответствующие
		// поля структуры ProjectTask
		if err := rows.Scan(&task.ID, &task.Name,
			&task.Description, &task.Priority,
			&task.IsDone, &task.ProjectID); err != nil {
			return nil, err
		}
		// Добавляем задачу в срез
		tasks = append(tasks, task)
	}
	return
}

// Метод для получения всех задач конкретного проекта
func (s *SQLiteRepository) GetProjectTasks(projectID int) (tasks []Task, err error) {
	// Запрос на получение всех задач у проекта с заданным id
	rows, err := s.db.Query(
		"SELECT * FROM tasks WHERE project_id = ?",
		projectID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Закрываем соединение

	for rows.Next() {
		var task Task
		var progID int
		// считываем данные из каждой строки, в соответствующие
		// поля структуры Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Description,
			&task.Priority, &task.IsDone, &progID); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return
}

// Метод для обновления задачи
func (s *SQLiteRepository) TaskDone(taskId int) error {
	if taskId == 0 { // Проверка на валидность
		return errors.New("invalid updated ID")
	}
	// Запрос на перевод задачи с указанным id
	// в состояние "выполнена"
	res, err := s.db.Exec(
		"UPDATE tasks SET is_done = ? WHERE id = ?", 1,
		taskId,
	)
	if err != nil {
		return err
	}

	// Проверяем, была ли обновлена задача
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// Если не была обновлена ни одна задача
	if rowsAffected == 0 {
		return ErrUpdateFailed
	}

	return nil
}
