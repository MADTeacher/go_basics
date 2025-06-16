package db

import "errors"

const dbName = "todo.db"

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

const ProjectTabelDefinition = `
CREATE TABLE IF NOT EXISTS projects(
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	name TEXT UNIQUE,
	description TEXT
);
`

const TaskTabelDefinition = `
CREATE TABLE IF NOT EXISTS tasks(
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	priority INTEGER NOT NULL,
	is_done BOOLEAN NOT NULL CHECK (is_done IN (0, 1)),
	project_id INTEGER not null references projects(id)
);
`
