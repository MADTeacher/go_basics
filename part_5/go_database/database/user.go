package database

import (
	"fmt"
	"strconv"
	"strings"
)

type AccessLevel = string

const (
	Student AccessLevel = "S"
	Teacher AccessLevel = "T"
	Admin   AccessLevel = "A"
)

func AcsessLevelFromString(s string) AccessLevel {
	switch strings.ToUpper(s) {
	case "T":
		return Teacher
	case "A":
		return Admin
	default:
		return Student
	}
}

type UserTableField = string

const (
	FieldID           UserTableField = "id"
	FieldNickname     UserTableField = "nickname"
	FieldYearOfBirth  UserTableField = "yearOfBirth"
	FieldEmail        UserTableField = "email"
	FieldPhone        UserTableField = "phone"
	FieldAccessLevel  UserTableField = "accessLevel"
	FieldPasswordHash UserTableField = "passwordHash"
)

type User struct {
	id           int
	nickname     string
	yearOfBirth  int
	email        string
	phone        string
	accessLevel  AccessLevel
	passwordHash string
}

func NewUser(
	id int,
	nickname string,
	yearOfBirth int,
	email string,
	phone string,
	accessLevel string,
	passwordHash string,
) *User {
	return &User{
		id:           id,
		nickname:     nickname,
		yearOfBirth:  yearOfBirth,
		email:        email,
		phone:        phone,
		accessLevel:  AcsessLevelFromString(accessLevel),
		passwordHash: passwordHash,
	}
}

func (u *User) ID() int              { return u.id }
func (u *User) Nickname() string     { return u.nickname }
func (u *User) YearOfBirth() int     { return u.yearOfBirth }
func (u *User) Email() string        { return u.email }
func (u *User) Phone() string        { return u.phone }
func (u *User) AccessLevel() string  { return u.accessLevel }
func (u *User) PasswordHash() string { return u.passwordHash }

func (u *User) Change(attr, value string) bool {
	switch attr {
	case FieldNickname:
		u.nickname = value
	case FieldYearOfBirth:
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		u.yearOfBirth = v
	case FieldEmail:
		u.email = value
	case FieldPhone:
		u.phone = value
	case FieldAccessLevel:
		u.accessLevel = AcsessLevelFromString(value)
	case FieldPasswordHash:
		u.passwordHash = value
	default:
		return false
	}
	return true
}

func (u *User) Check(attr, value string) bool {
	switch attr {
	case FieldID:
		v, err := strconv.Atoi(value)
		return err == nil && u.id == v
	case FieldNickname:
		return u.nickname == value
	case FieldYearOfBirth:
		v, err := strconv.Atoi(value)
		return err == nil && u.yearOfBirth == v
	case FieldEmail:
		return u.email == value
	case FieldPhone:
		return u.phone == value
	case FieldAccessLevel:
		return u.accessLevel == value
	case FieldPasswordHash:
		return u.passwordHash == value
	default:
		return false
	}
}

func (u *User) String() string {
	return fmt.Sprintf(
		"User(id: %d, nickname: %s, yearOfBirth: %d, "+
			"email: %s, phone: %s, acsessLevel: %s, passwordHash: %s)",
		u.id, u.nickname, u.yearOfBirth, u.email,
		u.phone, u.accessLevel, u.passwordHash,
	)
}
