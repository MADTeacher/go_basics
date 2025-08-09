package service

import (
	"encoding/json"
	"golang/todo-service/db"
	"io"
	"net/http"
	"strconv"
)

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	queryStr := r.URL.Query()

	idStr, ok := queryStr["id"]
	if !ok {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}
	if len(idStr) == 0 {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}

	id, err := strconv.Atoi(idStr[0])
	if err != nil {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}

	err = repository.DeleteTask(id)
	if err != nil {
		errorResponse(w, ErrorResponse{http.StatusNotFound,
			"Task not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	queryStr := r.URL.Query()

	idStr, ok := queryStr["projectID"]
	if !ok {
		tasks, err := repository.GetAllTasks()
		if err != nil || len(tasks) == 0 {
			errorResponse(w, ErrorResponse{http.StatusNotFound,
				"Tasks not found"})
			return
		}
		// возвращаем все имеющиеся задачи
		projectsList := TasksList{tasks}
		jsonGoodResponse, _ := json.Marshal(projectsList)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonGoodResponse)
		return
	}

	if len(idStr) == 0 {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}

	id, err := strconv.Atoi(idStr[0])
	if err != nil {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}

	tasks, err := repository.GetProjectTasks(id)
	if err != nil || len(tasks) == 0 {
		errorResponse(w, ErrorResponse{http.StatusNotFound,
			"Tasks not found"})
		return
	}
	// возвращаем все задачи конкретного проекта
	projectsList := TasksList{tasks}
	jsonGoodResponse, _ := json.Marshal(projectsList)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonGoodResponse)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var task db.Task
	reqBody, err := io.ReadAll(r.Body)
	if err == nil {
		json.Unmarshal(reqBody, &task)
		if task.Name == "" || task.ProjectID == 0 {
			errorResponse(w, ErrorResponse{http.StatusBadRequest,
				"Validation Failed"})
			return
		}

		task.ID = 0
		task, err := repository.AddTask(task)
		if err != nil {
			errorResponse(w, ErrorResponse{http.StatusBadRequest,
				"Validation Failed"})
			return
		}

		goodResponse := GoodResponse{
			Code:    201,
			Message: "Задача создана",
			ID:      task.ID,
		}
		jsonGoodResponse, _ := json.Marshal(goodResponse)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonGoodResponse)
		return
	}
	errorResponse(w, ErrorResponse{http.StatusBadRequest,
		"Validation Failed"})
}

func doneTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	queryStr := r.URL.Query()

	idStr, ok := queryStr["id"]
	if !ok {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}
	if len(idStr) == 0 {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}

	id, err := strconv.Atoi(idStr[0])
	if err != nil {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}

	err = repository.TaskDone(id)
	if err != nil {
		errorResponse(w, ErrorResponse{http.StatusNotFound,
			"Tasks not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
}
