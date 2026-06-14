package service

import (
	"encoding/json"
	"golang/todo-service/db"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)

	idStr, ok := params["id"]
	if !ok {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, ErrorResponse{http.StatusBadRequest,
			"Validation Failed"})
		return
	}

	err = repository.DeleteProject(id)
	if err != nil {
		errorResponse(w, ErrorResponse{http.StatusNotFound,
			"Project not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var project db.Project
	reqBody, err := io.ReadAll(r.Body)
	if err == nil {
		json.Unmarshal(reqBody, &project)
		if project.Name == "" {
			errorResponse(w, ErrorResponse{http.StatusBadRequest,
				"Validation Failed"})
			return
		}

		project.ID = 0
		project, err := repository.AddProject(project)
		if err != nil {
			errorResponse(w, ErrorResponse{http.StatusBadRequest,
				"Project with that name already exists"})
			return
		}

		goodResponse := GoodResponse{
			Code:    201,
			Message: "Проект создан",
			ID:      project.ID,
		}
		jsonGoodResponse, _ := json.Marshal(goodResponse)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonGoodResponse)
		return
	}
	errorResponse(w, ErrorResponse{http.StatusBadRequest,
		"Validation Failed"})
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	progects, err := repository.GetAllProjects()
	if err != nil || len(progects) == 0 {
		errorResponse(w, ErrorResponse{http.StatusNotFound,
			"Projects not found"})
		return
	}

	projectsList := ProjectsList{progects}
	jsonGoodResponse, _ := json.Marshal(projectsList)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonGoodResponse)
}
