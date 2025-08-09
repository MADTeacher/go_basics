package database

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type DBType int

const (
	SUAI DBType = iota
	UNECON
)

type Database struct {
	PathToSuaiDB   string
	PathToUneconDB string
	SuaiUsers      *Table[*User]
	UneconUsers    *Table[*User]
}

func NewDatabase(pathToSuaiDB, pathToUneconDB string) *Database {
	db := &Database{
		PathToSuaiDB:   pathToSuaiDB,
		PathToUneconDB: pathToUneconDB,
	}

	if _, err := os.Stat(pathToSuaiDB); os.IsNotExist(err) {
		dir := filepath.Dir(pathToSuaiDB)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
		}

		file, err := os.Create(pathToSuaiDB)
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		file.Close()
		db.SuaiUsers = NewTable[*User]("suai")
	} else {
		db.SuaiUsers = db.openTable(pathToSuaiDB, "suai")
	}

	if _, err := os.Stat(pathToUneconDB); os.IsNotExist(err) {
		dir := filepath.Dir(pathToUneconDB)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
		}

		file, err := os.Create(pathToUneconDB)
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		file.Close()
		db.UneconUsers = NewTable[*User]("unecon")
	} else {
		db.UneconUsers = db.openTable(pathToUneconDB, "unecon")
	}

	return db
}

func (db *Database) openTable(filePath, tableName string) *Table[*User] {
	table := NewTable[*User](tableName)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return table
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ",")
		if len(data) >= 7 {
			id, _ := strconv.Atoi(data[0])
			yearOfBirth, _ := strconv.Atoi(data[2])

			user := NewUser(
				id,
				data[1],
				yearOfBirth,
				data[3],
				data[4],
				data[5],
				data[6],
			)

			table.Insert(user)
		}
	}

	return table
}

func (db *Database) Save(dbType DBType) {
	var filePath string
	var users *Table[*User]

	switch dbType {
	case SUAI:
		filePath = db.PathToSuaiDB
		users = db.SuaiUsers
	case UNECON:
		filePath = db.PathToUneconDB
		users = db.UneconUsers
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error opening file for writing:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	users.ForEach(func(user *User) {
		line := fmt.Sprintf("%d,%s,%d,%s,%s,%s,%s\n",
			user.ID(),
			user.Nickname(),
			user.YearOfBirth(),
			user.Email(),
			user.Phone(),
			user.AccessLevel(),
			user.PasswordHash(),
		)
		writer.WriteString(line)
	})
	writer.Flush()
}

func (db *Database) Insert(user *User, dbType DBType) bool {
	var filePath string
	var isOk bool

	switch dbType {
	case SUAI:
		filePath = db.PathToSuaiDB
		isOk = db.SuaiUsers.Insert(user)
	case UNECON:
		filePath = db.PathToUneconDB
		isOk = db.UneconUsers.Insert(user)
	}

	if !isOk {
		return false
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for appending:", err)
		return false
	}
	defer file.Close()

	line := fmt.Sprintf("%d,%s,%d,%s,%s,%s,%s\n",
		user.ID(),
		user.Nickname(),
		user.YearOfBirth(),
		user.Email(),
		user.Phone(),
		user.AccessLevel(),
		user.PasswordHash(),
	)

	if _, err := file.WriteString(line); err != nil {
		fmt.Println("Error writing to file:", err)
		return false
	}

	return true
}

func (db *Database) Selection(dbType DBType,
	attribute, value string) *Table[*User] {
	switch dbType {
	case SUAI:
		return db.SuaiUsers.Selection(attribute, value)
	case UNECON:
		return db.UneconUsers.Selection(attribute, value)
	default:
		return NewTable[*User]("empty")
	}
}

func (db *Database) Intersect(attribute, value string) *Table[*User] {
	return db.SuaiUsers.Intersect(attribute, value, db.UneconUsers)
}

func (db *Database) Union() *Table[*User] {
	return db.SuaiUsers.Union(db.UneconUsers)
}

func (db *Database) Remove(id string, dbType DBType) {
	switch dbType {
	case SUAI:
		db.SuaiUsers.Remove(id)
		fmt.Println(db.SuaiUsers)
	case UNECON:
		db.UneconUsers.Remove(id)
		fmt.Println(db.UneconUsers)
	}
}

func (db *Database) ShowDB(dbType DBType) {
	switch dbType {
	case SUAI:
		fmt.Println(db.SuaiUsers)
	case UNECON:
		fmt.Println(db.UneconUsers)
	}
}
