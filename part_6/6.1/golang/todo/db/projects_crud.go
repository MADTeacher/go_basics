package db

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

// Метод для добавления проекта в базу данных
func (s *SQLiteRepository) AddProject(project Project) (*Project, error) {
	res, err := s.db.Exec( // Запрос на добавление проекта
		// Используем экранирование данных для предотвращения SQL-инъекции
		// данные из переменных project.Name и project.Description
		// будут вставлены в запрос в место знаков вопроса
		// в пордке их перечисления
		"INSERT INTO projects(name, description) values(?,?)",
		project.Name, project.Description,
	)
	if err != nil { // Если произошла ошибка
		var sqliteErr sqlite3.Error

		// Если такой проект уже существует
		if errors.As(err, &sqliteErr) {
			if errors.Is(
				sqliteErr.ExtendedCode,
				sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate // Возвращаем ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId() // Получаем ID
	if err != nil {
		return nil, err
	}
	project.ID = int(id) // Устанавливаем ID проекту

	return &project, nil
}

// Метод для удаления проекта из базы данных
func (s *SQLiteRepository) DeleteProject(projectID int) error {
	// Запрос на удаление проекта из таблицы projects
	s.db.Exec("DELETE FROM projects WHERE id = ?", projectID)
	// Запрос на удаление всех задач, связанных с проектом
	res, err := s.db.Exec(
		"DELETE FROM tasks WHERE project_id = ?",
		projectID,
	)
	if err != nil {
		return err
	}

	// Проверяем, были ли удалены задачи
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// Если не было удалено ни одной задачи
	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}

// Метод для получения всех проектов
func (s *SQLiteRepository) GetAllProjects() ([]Project, error) {
	// Запрос на получение всех проектов
	rows, err := s.db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Закрываем соединение

	// Перебираем результаты запроса и добавляем в []Project
	var projects []Project
	for rows.Next() {
		var project Project
		// считываем данные из каждой строки, в соответствующие
		// поля структуры Project
		if err := rows.Scan(&project.ID, &project.Name,
			&project.Description); err != nil {
			return nil, err
		}
		// Добавляем проект в срез
		projects = append(projects, project)
	}
	return projects, nil
}
