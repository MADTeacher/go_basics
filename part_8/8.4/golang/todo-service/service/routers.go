package service

import (
	"encoding/json"
	"fmt"
	"golang/todo-service/db"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string           // имя функции обработчика
	Method      string           // тип HTTP-сообщения
	Pattern     string           // шаблон пути
	HandlerFunc http.HandlerFunc // ссылка на функцию обработчик
	// сигнатура функции должна быть func(ResponseWriter, *Request)
}

var repository *db.SQLiteRepository

func NewRouter(rep *db.SQLiteRepository) *mux.Router {
	repository = rep
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		//оборачиваем функцию-обработчик в логгер
		handler = Logger(route.HandlerFunc, route.Name)
		// 	handler := route.HandlerFunc
		// 	// 	// добавляем новый обработчик
		router. //HandleFunc(route.Pattern, handler).Methods(route.Method)
			Methods(route.Method). // тип HTTP-сообщения
			Path(route.Pattern).   // шаблон пути
			Name(route.Name).      // имя функции обработчика
			Handler(handler)       // ссылка на функцию обработчик
	}
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wold in REST API style!!!")
}

var routes = []Route{
	{ // домашняя страница
		"home",
		http.MethodGet,
		"/api/v1/todo",
		home,
	},
	{ // удаление проекта по id
		"deleteProject",
		http.MethodDelete,
		"/api/v1/todo/project/del/{id}",
		deleteProject,
	},
	{ // добавление проекта
		"addProject",
		http.MethodPost,
		"/api/v1/todo/project",
		addProject,
	},
	{ // получить все проекты
		"getProjects",
		http.MethodGet,
		"/api/v1/todo/projects",
		getProjects,
	},
	{ // удаление задачи
		"deleteTask",
		http.MethodDelete,
		"/api/v1/todo/task",
		deleteTask,
	},
	{ // получение всех задач или конкретного проекта
		"getTask",
		http.MethodGet,
		"/api/v1/todo/task",
		getTask,
	},
	{ // добавить задачу
		"addTask",
		http.MethodPost,
		"/api/v1/todo/task",
		addTask,
	},
	{ // изменение статуса задачи на «Выполнено»
		"doneTask",
		http.MethodPut,
		"/api/v1/todo/task",
		doneTask,
	},
}

func errorResponse(w http.ResponseWriter, err ErrorResponse) {
	jsonResponse, _ := json.Marshal(err)
	w.WriteHeader(err.Code)
	w.Write(jsonResponse)
}
