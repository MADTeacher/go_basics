package menu

import (
	"bufio"
	"fmt"
	"golang/todo/db"
	"os"
)

func addTask(rep *db.SQLiteRepository) {
	task := db.Task{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input project ID: ")
	projectID, err := getIntValueFromStd(reader)
	if err != nil {
		printNotValidData()
		return
	}

	fmt.Print("Input task name: ")
	name, _ := getStringValueFromStd(reader)

	fmt.Print("Input description task: ")
	desc, _ := getStringValueFromStd(reader)

	fmt.Print("Input priority task: ")
	priority, err := getIntValueFromStd(reader)
	if err != nil {
		printNotValidData()
		return
	}

	task.Name = name
	task.Description = desc
	task.Priority = uint8(priority)
	if task.Name != "" && task.Description != "" {
		task, err := rep.AddTask(task, projectID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\nAdded task: %+v\n", *task)
	} else {
		printNotValidData()
	}
}

func deleteTaskByID(rep *db.SQLiteRepository) {
	fmt.Print("Input ID for deleting task: ")
	id, err := getIntValueFromStd(bufio.NewReader(os.Stdin))
	if err != nil {
		printNotValidData()
		return
	}

	err = rep.DeleteTask(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Task deleted")
}

func getAllTasks(rep *db.SQLiteRepository) {
	tasks, err := rep.GetAllTasks()
	if err != nil {
		printNotValidData()
		return
	}
	if len(tasks) == 0 {
		fmt.Println("You don't have any task")
	} else {
		fmt.Println("You current tasks: ")
		for _, it := range tasks {
			fmt.Printf("TaskID: %v || Name: %v || Desc: %v || Priority: %v || IsDone: %v || ProjID: %v\n",
				it.ID, it.Name, it.Description, it.Priority, it.IsDone, it.ProjectID)
		}
	}
}

func getAllProjectTasks(rep *db.SQLiteRepository) {
	fmt.Print("Input ID for project: ")

	id, err := getIntValueFromStd(bufio.NewReader(os.Stdin))
	if err != nil {
		printNotValidData()
		return
	}

	tasks, err := rep.GetProjectTasks(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("You don't have any task")
	} else {
		fmt.Printf("Project with ID = %d have next tasks:\n", id)
		for _, it := range tasks {
			fmt.Printf("TaskID: %v || Name: %v || Desc: %v || Priority: %v || IsDone: %v\n",
				it.ID, it.Name, it.Description, it.Priority, it.IsDone)
		}
	}
}

func doneTask(rep *db.SQLiteRepository) {
	fmt.Print("Input task ID: ")

	id, err := getIntValueFromStd(bufio.NewReader(os.Stdin))
	if err != nil {
		printNotValidData()
		return
	}

	err = rep.TaskDone(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Congratulations! Task done!")
}
