package db

import (
	"errors"

	"github.com/mattn/go-sqlite3"
)

func (s *SQLiteRepository) AddProject(project Project) (*Project, error) {
	res, err := s.db.Exec("INSERT INTO projects(name, description) values(?,?)",
		project.Name, project.Description)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	project.ID = int(id)

	return &project, nil
}

func (s *SQLiteRepository) DeleteProject(projectID int) error {
	s.db.Exec("DELETE FROM projects WHERE id = ?", projectID)
	res, err := s.db.Exec("DELETE FROM tasks WHERE project_id = ?", projectID)
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

func (s *SQLiteRepository) GetAllProjects() ([]Project, error) {
	rows, err := s.db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}
