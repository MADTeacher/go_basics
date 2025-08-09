package database

type IAttribute interface {
	Check(attribute, value string) bool
	Change(attribute, value string) bool

	String() string
}
