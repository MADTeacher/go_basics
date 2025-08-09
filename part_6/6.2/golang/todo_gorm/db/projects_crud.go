package db

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

// Метод для добавления проекта в базу данных
func (r *SQLiteRepository) AddProject(project Project) (*Project, error) {
	tx := r.db.Create(&project)
	if tx.Error != nil {
		var sqliteErr sqlite3.Error
		if errors.As(tx.Error, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode,
				sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, tx.Error
	}

	return &project, nil
}

// Метод для удаления проекта из базы данных
func (r *SQLiteRepository) DeleteProject(projectID int) error {
	tx := r.db.Delete(&Project{ID: projectID})
	if tx.Error != nil {
		return tx.Error
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return nil
}

// Метод для получения всех проектов
func (r *SQLiteRepository) GetAllProjects() ([]Project, error) {
	var projects []Project
	tx := r.db.Find(&projects)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, ErrNotExists
	}

	return projects, nil
}
