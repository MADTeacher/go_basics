package db

type Project struct {
	ID          int    `json:"id" gorm:"primary_key;autoIncrement:true;not null"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
}

type Task struct {
	ID          int      `json:"id" gorm:"primary_key;autoIncrement;not null"`
	Name        string   `json:"name" gorm:"not null"`
	Description string   `json:"description" gorm:"not null"`
	Priority    uint8    `json:"priority" gorm:"not null"`
	IsDone      bool     `json:"isDone" gorm:"not null"`
	ProjectID   int      `json:"projectID" gorm:"not null"`
	Project     *Project `gorm:"foreignKey:ProjectID;references:ID"`
}
