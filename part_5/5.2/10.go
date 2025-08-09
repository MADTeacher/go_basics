package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	countDirs := 0
	countFiles := 0
	filepath.WalkDir("C:\\Go", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			countFiles++
		} else {
			countDirs++
		}
		return nil
	})

	fmt.Printf("Amount of files: %d\n", countFiles)
	fmt.Printf("Amount of directories: %d\n", countDirs)
}
