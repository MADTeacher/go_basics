package db

type Project struct {
	ID          int
	Name        string
	Description string
}

type Task struct {
	ID          int
	Name        string
	Description string
	Priority    uint8
	IsDone      bool
}

type ProjectTask struct {
	Task
	ProjectID int
}
