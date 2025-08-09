package menu

import (
	"bufio"
	"fmt"
	"golang/todo/db"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dixonwille/wmenu"
)

func CreateMenu(rep *db.SQLiteRepository) {
	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(rep, opts); return nil })

	menu.Option("Add a new Project", 0, false, nil)
	menu.Option("Delete a Project by ID", 1, false, nil)
	menu.Option("Get all Projects", 2, false, nil)
	menu.Option("Add a Task", 3, false, nil)
	menu.Option("Get all Tasks", 4, true, nil) // выбор по умолчанию
	menu.Option("Get all Project tasks", 5, false, nil)
	menu.Option("Done a Task by ID", 6, false, nil)
	menu.Option("Delete a Task by ID", 7, false, nil)
	menu.Option("Quit Application", 8, false, nil)
	menuerr := menu.Run()
	fmt.Println()
	fmt.Println("---------------------------------")
	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

func handleFunc(rep *db.SQLiteRepository, opts []wmenu.Opt) {
	switch opts[0].Value {
	case 0:
		fmt.Println("Adding a new Project")
		addProject(rep)
	case 1:
		fmt.Println("Deleting a Project by ID")
		deleteProjectByID(rep)
	case 2:
		fmt.Println("Getting all Projects")
		getAllProjects(rep)
	case 3:
		fmt.Println("Adding a new Task")
		addTask(rep)
	case 4:
		fmt.Println("Getting all Tasks")
		getAllTasks(rep)
	case 5:
		fmt.Println("Getting all Project tasks by ProjectID")
		getAllProjectTasks(rep)
	case 6:
		fmt.Println("Doing a Task by ID")
		doneTask(rep)
	case 7:
		fmt.Println("Deleting a Task by ID")
		deleteTaskByID(rep)
	case 8:
		fmt.Println("See you later!!!")
		os.Exit(0)
	}
}

func printNotValidData() {
	fmt.Println("Data is not valid!!!")
}

func getIntValueFromStd(reader *bufio.Reader) (int, error) {
	tempID, _, _ := reader.ReadLine()
	idStr := strings.TrimSuffix(string(tempID), "\n")
	idProj, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return idProj, nil
}

func getStringValueFromStd(reader *bufio.Reader) (string, error) {
	data, err := reader.ReadString('\n')
	data = strings.TrimSuffix(data, "\r\n")
	if err != nil {
		return "", err
	}
	return data, nil
}
