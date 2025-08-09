package db

type Project struct {
	ID          int    `gorm:"primary_key;autoIncrement:true;not null"`
	Name        string `gorm:"unique;not null"`
	Description string
}

type Task struct {
	ID          int    `gorm:"primary_key;autoIncrement;not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Priority    uint8  `gorm:"not null"`
	IsDone      bool   `gorm:"not null"`
}

type ProjectTask struct {
	Task
	ProjectID int      `gorm:"not null"`
	Project   *Project `gorm:"foreignKey:ProjectID;references:ID"`
}
