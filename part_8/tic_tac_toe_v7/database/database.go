package database

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	db *gorm.DB // заменили на *gorm.DB
}

func NewSQLiteRepository() (*SQLiteRepository, error) {
	// Создаем репозиторий
	repository := &SQLiteRepository{}

	// Проверяем существование файла базы данных
	dbExists := true
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		dbExists = false
	}

	// Открываем соединение с базой данных
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Сохраняем соединение в репозитории
	repository.db = db

	// Если база данных только что создана, выполняем миграцию
	if !dbExists {
		fmt.Println("Creating new database schema")
		if err := db.AutoMigrate(&PlayerFinishGame{}); err != nil {
			return nil, fmt.Errorf("failed to migrate database: %w", err)
		}
	} else {
		fmt.Println("Using existing database")
	}

	return repository, nil
}
