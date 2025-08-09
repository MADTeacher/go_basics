package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   uint
	Name string
	Age  uint8
}

type UserDB struct {
	users []User
}

func (db *UserDB) FindUserWithID(id uint) (*User, error) {
	for _, it := range db.users {
		if it.ID == id {
			return &it, nil
		}
	}
	return nil, fmt.Errorf("User with id=%v not found", id)
}

func (db *UserDB) SetUserAge(id uint, age uint8) error {
	for _, it := range db.users {
		if it.ID == id {
			it.Age = age
			return nil
		}
	}
	return fmt.Errorf("User with id=%v not found", id)
}

func (db *UserDB) AddUser(user User) error {
	for _, it := range db.users {
		if it.ID == user.ID {
			return fmt.Errorf("User with id=%v already exist", user.ID)
		}
	}
	db.users = append(db.users, user)
	return nil
}

func (db *UserDB) DeleteUserWithName(name string) error {
	index := -1
	for idx, it := range db.users {
		if it.Name == name {
			index = idx
			break
		}
	}
	if index < 0 {
		return fmt.Errorf("User %s not found", name)
	}
	db.users = append(db.users[:index], db.users[index+1:]...)
	return nil
}

// функции для работы с БД и организации цепочек ошибок
func findUserWithID(db *UserDB, id uint) (*User, error) {
	user, err := db.FindUserWithID(id)
	if err != nil {
		// вернется ошибка из 2-х вложений
		return nil, fmt.Errorf("DataBase error: %w", err)
	}
	return user, nil
}

func setUserAge(db *UserDB, id uint, age uint8) error {
	err := db.SetUserAge(id, age)
	if err != nil {
		// вернется ошибка из 2-х вложений
		return fmt.Errorf("DataBase user data error: %w", err)
	}
	return nil
}

func addUser(db *UserDB, user User) error {
	err := db.AddUser(user)
	if err != nil {
		// вернется ошибка из 2-х вложений
		return fmt.Errorf("Add user error: %w", err)
	}
	return nil
}

func main() {
	dbUsers := UserDB{
		[]User{
			{2, "Alex", 28},
			{1, "Max", 23},
			{10, "German", 35},
			{6, "Oleg", 19},
		},
	}

	user, err := findUserWithID(&dbUsers, 2)
	if err != nil {
		fmt.Println(err)
		// распаковка трассировки ошибок
		fmt.Println(errors.Unwrap(err))
	}
	fmt.Printf("%+v\n", user) // &{ID:2 Name:Alex Age:28}

	user, err = findUserWithID(&dbUsers, 12)
	if err != nil {
		fmt.Println(err) // DataBase error: User with id=12 not found
		// распаковка трассировки ошибок
		fmt.Println(errors.Unwrap(err)) // User with id=12 not found
	}
	fmt.Printf("%+v\n", user) // <nil>

	newErr := addUser(&dbUsers, User{6, "Li", 45})
	if newErr != nil {
		fmt.Println(newErr) // Add user error: User with id=6 already exist
		// распаковка трассировки ошибок
		fmt.Println(errors.Unwrap(newErr)) // User with id=6 already exist
	}
}
