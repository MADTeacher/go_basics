package db

import (
	_ "database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	db *gorm.DB // заменили на *gorm.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	var db *gorm.DB
	rep := &SQLiteRepository{}
	// Если база данных не существует, то создаем ее
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		// Отквываем соединение с базой данных
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DB isn't exist")
		// Создаем таблицы
		db.AutoMigrate(&Project{}, &ProjectTask{})
		// Заполняем БД значениями по умолчанию
		rep.db = db
		putDefaultValuesToDB(rep)
	} else {
		// Отквываем соединение с базой данных
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		rep.db = db
		fmt.Println("DB already exists")
	}

	return rep
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

func (r *SQLiteRepository) Close() {

}
