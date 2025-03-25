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
	db *gorm.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	var db *gorm.DB

	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DB isn't exist")
		db.AutoMigrate(&Project{}, &ProjectTask{})
		putDefaultValuesToDB(db)
	} else {
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DB already exists")
	}

	return &SQLiteRepository{
		db: db,
	}
}

func putDefaultValuesToDB(db *gorm.DB) {
	firstProject := Project{
		Name:        "Go",
		Description: "Roadmap for learning Go",
	}
	secondProject := Project{
		Name:        "One Year",
		Description: "Tasks for the year",
	}
	db.Create(&firstProject)
	db.Create(&secondProject)
	db.Create(&ProjectTask{
		Task: Task{
			Name:        "Variable",
			Description: "Learning Go build-in variables",
			Priority:    1,
		},
		Project: &firstProject,
	})
	db.Create(&ProjectTask{
		Task: Task{
			Name:        "Struct",
			Description: "Learning use struct in OOP code",
			Priority:    3,
		},
		Project: &firstProject,
	})
	db.Create(&ProjectTask{
		Task: Task{
			Name:        "Goroutine",
			Description: "Learning concurrent programming",
			Priority:    5,
		},
		Project: &firstProject,
	})
	db.Create(&ProjectTask{
		Task: Task{
			Name:        "DataBase",
			Description: "How write app with db",
			Priority:    1,
		},
		Project: &firstProject,
	})
	db.Create(&ProjectTask{
		Task: Task{
			Name:        "PhD",
			Description: "Ph.D. in Technical Sciences",
			Priority:    5,
		},
		Project: &secondProject,
	})
	db.Create(&ProjectTask{
		Task: Task{
			Name:        "Losing weight",
			Description: "Exercise and eat less chocolate",
			Priority:    2,
		},
		Project: &secondProject,
	})
	db.Create(&ProjectTask{
		Task: Task{
			Name:        "Пафос и превозмогание",
			Description: "10к подписчиков на канале",
			Priority:    2,
		},
		Project: &secondProject,
	})
}

func (r *SQLiteRepository) Close() {

}
