package menu

import (
	"bufio"
	"fmt"
	"golang/todo/db"
	"os"
)

// Функция для добавления нового проекта
func addProject(rep *db.SQLiteRepository) {
	project := db.Project{}             // Создаем новый проект
	reader := bufio.NewReader(os.Stdin) // Создаем поток ввода

	fmt.Print("Input project name: ")
	name, _ := getStringValueFromStd(reader)

	fmt.Print("Input description project: ")
	desc, _ := getStringValueFromStd(reader)

	project.Name = name
	project.Description = desc
	// Если название и описание проекта не пустые
	if project.Name != "" && project.Description != "" {
		project, err := rep.AddProject(project) // Добавляем проект
		if err != nil {
			fmt.Println(err)
		} else {
			// Выводим информацию о добавленном проекте
			fmt.Printf("\nAdded project: %+v\n", *project)
		}
	} else {
		// Выводим сообщение об ошибке
		printNotValidData()
	}
}

// Функция для удаления проекта
func deleteProjectByID(rep *db.SQLiteRepository) {
	fmt.Print("Input ID for deleting project: ")
	id, err := getIntValueFromStd(bufio.NewReader(os.Stdin))
	if err != nil {
		printNotValidData()
		return
	}

	err = rep.DeleteProject(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Project deleted")
}

// Функция для получения всех проектов
func getAllProjects(rep *db.SQLiteRepository) {
	progects, err := rep.GetAllProjects()
	if err != nil {
		printNotValidData()
		return
	}
	if len(progects) == 0 {
		fmt.Println("You don't have any project")
	} else {
		fmt.Println("You current projects:")
		for _, it := range progects {
			fmt.Printf("ProjectID: %v || Name: %v || Desc: %v\n",
				it.ID, it.Name, it.Description)
		}
	}
}
