package main

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "todo.db"

var ErrNotExists = errors.New("row not exists")

type Project struct {
	ID          int    `gorm:"primary_key;autoIncrement:true;not null"`
	Name        string `gorm:"unique;not null"`
	Description string
}

type ProjectTask struct {
	ID          int      `gorm:"primary_key;autoIncrement;not null"`
	Name        string   `gorm:"not null"`
	Description string   `gorm:"not null"`
	Priority    uint8    `gorm:"not null"`
	IsDone      bool     `gorm:"not null"`
	ProjectID   int      `gorm:"not null"`
	Project     *Project `gorm:"foreignKey:ProjectID;references:ID"`
}

func connectionToBD(pathToDB string) (db *gorm.DB, err error) {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		fmt.Println("DB isn't exist")
		db.AutoMigrate(&Project{}, &ProjectTask{})
		putDefaultValuesToDB(db)
	} else {
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		fmt.Println("DB already exists")
	}
	return
}

func putDefaultValuesToDB(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error { // начало транзакции
		firstProject := Project{
			Name:        "Go",
			Description: "Roadmap for learning Go",
		}
		secondProject := Project{
			Name:        "One Year",
			Description: "Tasks for the year",
		}

		tx.Create(&firstProject)

		if err := tx.Create(&secondProject).Error; err != nil { //проверяем на наличие ошибок при записи
			return err // вызываем отмену транзакции
		}

		tx.Create(&ProjectTask{
			Name:        "Variable",
			Description: "Learning Go build-in variables",
			Priority:    1,
			Project:     &firstProject,
		})

		if err := tx.Create(&ProjectTask{ //проверяем на наличие ошибок при записи
			Name:        "Struct",
			Description: "Learning use struct in OOP code",
			Priority:    3,
			Project:     &firstProject,
		}).Error; err != nil {
			return err // вызываем отмену транзакции
		}

		tx.Create(&ProjectTask{
			Name:        "Goroutine",
			Description: "Learning concurrent programming",
			Priority:    5,
			Project:     &firstProject,
		})
		tx.Create(&ProjectTask{
			Name:        "DataBase",
			Description: "How write app with db",
			Priority:    1,
			Project:     &firstProject,
		})
		tx.Create(&ProjectTask{
			Name:        "PhD",
			Description: "Ph.D. in Technical Sciences",
			Priority:    5,
			Project:     &secondProject,
		})
		tx.Create(&ProjectTask{
			Name:        "Losing weight",
			Description: "Exercise and eat less chocolate",
			Priority:    2,
			Project:     &secondProject,
		})
		tx.Create(&ProjectTask{
			Name:        "Пафос и превозмогание",
			Description: "10к подписчиков на канале",
			Priority:    2,
			Project:     &secondProject,
		})
		return nil
	})
}

func GetAllProjects(db *gorm.DB) ([]Project, error) {
	var projects []Project
	tx := db.Find(&projects)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, ErrNotExists
	}

	return projects, nil
}

func GetAllTasks(db *gorm.DB) (tasks []ProjectTask, err error) {
	tx := db.Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, ErrNotExists
	}

	return
}

func printAllTasks(db *gorm.DB) {
	tasks, err := GetAllTasks(db)
	if err != nil {
		return
	}
	fmt.Println("*********Tasks*********")
	for _, it := range tasks {
		fmt.Printf("TaskID: %v || Name: %v || Priority: %v || IsDone: %v || ProjID: %v\n",
			it.ID, it.Name, it.Priority, it.IsDone, it.ProjectID)
	}
}

func printAllProjects(db *gorm.DB) {
	progects, err := GetAllProjects(db)
	if err != nil {
		return
	}
	fmt.Println("*********Projects*********")
	for _, it := range progects {
		fmt.Printf("ProjectID: %v || Name: %v || Desc: %v\n",
			it.ID, it.Name, it.Description)
	}
}

func printAllProjectAndTask(db *gorm.DB) {
	printAllProjects(db)
	printAllTasks(db)
}

func GetAllProjectTasks(db *gorm.DB, projectID int) (tasks []ProjectTask, err error) {
	tx := db.Where("project_id", projectID).Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, ErrNotExists
	}

	return
}

func printAllProjectTasks(db *gorm.DB, projectID int) {
	tasks, err := GetAllProjectTasks(db, projectID)
	if err != nil {
		return
	}
	fmt.Println("*********Tasks*********")
	for _, it := range tasks {
		fmt.Printf("TaskID: %v || Name: %v || Priority: %v || IsDone: %v || ProjID: %v\n",
			it.ID, it.Name, it.Priority, it.IsDone, it.ProjectID)
	}
}

func main() {
	db, _ := gorm.Open(sqlite.Open(dbName), &gorm.Config{DryRun: true})

	var tasks []ProjectTask
	stmt := db.Where("project_id", 1).Find(&tasks).Statement
	fmt.Println(stmt.SQL.String())

	stmt = db.Create(&ProjectTask{
		Name:        "Пафос и превозмогание",
		Description: "10к подписчиков на канале",
		Priority:    2,
		ProjectID:   2,
	}).Statement
	fmt.Println(stmt.SQL.String())

	firstProject := Project{
		Name:        "Go",
		Description: "Roadmap for learning Go",
	}

	stmt = db.Create(&firstProject).Statement
	fmt.Println(stmt.SQL.String())
}

func CreateProjects(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&Project{Name: "Oo"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Project{Name: "^_^"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
