package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type SQLiteRepository struct {
	db *sql.DB
}

func createDB(pathToDB string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(ProjectTabelDefinition)
	db.Exec(TaskTabelDefinition)
	return db
}

func NewSQLiteRepository() *SQLiteRepository {
	var db *sql.DB

	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		db = createDB(dbName)
		fmt.Println("DB isn't exist")
		putDefaultValuesToDB(&SQLiteRepository{
			db: db,
		})
	} else {
		db, err = sql.Open("sqlite3", dbName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DB already exists")
	}

	return &SQLiteRepository{
		db: db,
	}
}

func putDefaultValuesToDB(rep *SQLiteRepository) {
	firstProject, _ := rep.AddProject(Project{
		Name:        "Go",
		Description: "Roadmap for learning Go",
	})
	secondProject, _ := rep.AddProject(Project{
		Name:        "One Year",
		Description: "Tasks for the year",
	})
	rep.AddTask(Task{
		Name:        "Variable",
		Description: "Learning Go build-in variables",
		Priority:    1,
	}, firstProject.ID)
	rep.AddTask(Task{
		Name:        "Struct",
		Description: "Learning use struct in OOP code",
		Priority:    3,
	}, firstProject.ID)
	rep.AddTask(Task{
		Name:        "Goroutine",
		Description: "Learning concurrent programming",
		Priority:    5,
	}, firstProject.ID)
	rep.AddTask(Task{
		Name:        "DataBase",
		Description: "How write app with db",
		Priority:    1,
	}, firstProject.ID)
	rep.AddTask(Task{
		Name:        "PhD",
		Description: "Ph.D. in Technical Sciences",
		Priority:    5,
	}, secondProject.ID)
	rep.AddTask(Task{
		Name:        "Losing weight",
		Description: "Exercise and eat less chocolate",
		Priority:    2,
	}, secondProject.ID)
	rep.AddTask(Task{
		Name:        "Пафос и превозмогание",
		Description: "10к подписчиков на канале",
		Priority:    2,
	}, secondProject.ID)
}

func (s *SQLiteRepository) Close() {
	s.db.Close()
}
