package main

import (
	"bufio"
	"fmt"
	db "go_database/database"
	"os"
	"strconv"
	"strings"
)

type Menu struct {
	database *db.Database
	reader   *bufio.Reader
}

func NewMenu(db *db.Database) *Menu {
	return &Menu{
		database: db,
		reader:   bufio.NewReader(os.Stdin),
	}
}

func (m *Menu) Loop() {
	for {
		m.PrintMenu()
		input, _ := m.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		fmt.Println(strings.Repeat("x", 25))

		switch input {
		case "1":
			m.AddUser()
		case "2":
			m.RemoveUser()
		case "3":
			m.ChangeUser()
		case "4":
			m.ShowUsers()
		case "5":
			m.Intersect()
		case "6":
			m.Union()
		case "7":
			m.SaveAndExit()
			return
		case "8":
			return
		}
	}
}

func (m *Menu) PrintMenu() {
	fmt.Println("1. Add User")
	fmt.Println("2. Remove User")
	fmt.Println("3. Change User")
	fmt.Println("4. Show Users")
	fmt.Println("5. Intersect Table")
	fmt.Println("6. Union Table")
	fmt.Println("7. Save and Exit")
	fmt.Println("8. Exit")
}

func (m *Menu) readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := m.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (m *Menu) AddUser() {
	try := func() bool {
		idStr := m.readInput("Enter id: ")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID format")
			return false
		}

		nickname := m.readInput("Enter nickname: ")

		yearStr := m.readInput("Enter year of birth: ")
		yearOfBirth, err := strconv.Atoi(yearStr)
		if err != nil {
			fmt.Println("Invalid year format")
			return false
		}

		email := m.readInput("Enter email: ")
		phone := m.readInput("Enter phone: ")
		passwordHash := m.readInput("Enter password hash: ")
		accessLevel := m.readInput("Access level (A - Admin, " +
			"T - Teacher, S - Student): ")

		user := db.NewUser(
			id,
			nickname,
			yearOfBirth,
			email,
			phone,
			accessLevel,
			passwordHash,
		)

		dbTypeStr := m.readInput("Add user to DB (S - SUAI, U - Unecon): ")
		dbTypeStr = strings.ToUpper(dbTypeStr)

		switch dbTypeStr {
		case "S":
			m.database.Insert(user, db.SUAI)
			m.database.ShowDB(db.SUAI)
		case "U":
			m.database.Insert(user, db.UNECON)
			m.database.ShowDB(db.UNECON)
		default:
			fmt.Println("(ノ-_-)ノ ミ ┴┴")
			return false
		}

		return true
	}

	if !try() {
		fmt.Println("WTF!!!!")
	}
}

func (m *Menu) RemoveUser() {
	dbTypeStr := m.readInput("Select DB (S - SUAI, U - Unecon): ")
	dbTypeStr = strings.ToUpper(dbTypeStr)

	var dbType db.DBType

	switch dbTypeStr {
	case "S":
		dbType = db.SUAI
		m.database.ShowDB(db.SUAI)
	case "U":
		dbType = db.UNECON
		m.database.ShowDB(db.UNECON)
	default:
		fmt.Println("(ノ-_-)ノ ミ ┴┴")
		return
	}

	id := m.readInput("Enter id: ")
	m.database.Remove(id, dbType)
	m.database.ShowDB(dbType)
}

func (m *Menu) ChangeUser() {
	dbTypeStr := m.readInput("Select DB (S - SUAI, U - Unecon): ")
	dbTypeStr = strings.ToUpper(dbTypeStr)

	var dbType db.DBType

	switch dbTypeStr {
	case "S":
		dbType = db.SUAI
		m.database.ShowDB(db.SUAI)
	case "U":
		dbType = db.UNECON
		m.database.ShowDB(db.UNECON)
	default:
		fmt.Println("(ノ-_-)ノ ミ ┴┴")
		return
	}

	id := m.readInput("Enter id: ")

	userTable := m.database.Selection(dbType, "id", id)
	user := userTable.First()
	if user == nil {
		fmt.Println("User not found!")
		return
	}

	field := m.readInput("Enter field: ")
	newValue := m.readInput("Enter new value: ")
	user.Change(field, newValue)

	m.database.ShowDB(dbType)
}

func (m *Menu) ShowUsers() {
	dbTypeStr := m.readInput("Select DB (S - SUAI, U - Unecon): ")
	dbTypeStr = strings.ToUpper(dbTypeStr)

	switch dbTypeStr {
	case "S":
		m.database.ShowDB(db.SUAI)
	case "U":
		m.database.ShowDB(db.UNECON)
	default:
		fmt.Println("(ノ-_-)ノ ミ ┴┴")
	}
}

func (m *Menu) Intersect() {
	field := m.readInput("Enter field: ")
	value := m.readInput("Enter intersect value: ")
	result := m.database.Intersect(field, value)
	fmt.Println(result)
}

func (m *Menu) Union() {
	result := m.database.Union()
	fmt.Println(result)
}

func (m *Menu) SaveAndExit() {
	m.database.Save(db.SUAI)
	m.database.Save(db.UNECON)
}
