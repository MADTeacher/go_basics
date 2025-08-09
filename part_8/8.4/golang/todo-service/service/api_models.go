package service

import "golang/todo-service/db"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GoodResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ID      int    `json:"id"`
}

type ProjectsList struct {
	Items []db.Project `json:"items,omitempty"`
}

type TasksList struct {
	Items []db.Task `json:"items,omitempty"`
}
