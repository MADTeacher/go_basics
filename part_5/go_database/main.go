package main

import (
	db "go_database/database"
	"path/filepath"
)

func main() {
	db := db.NewDatabase(
		filepath.Join(".", "suai.txt"),
		filepath.Join(".", "unecon.txt"),
	)

	menu := NewMenu(db)
	menu.Loop()
}
