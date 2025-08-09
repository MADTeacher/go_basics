package jsonstore

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type JSONFile struct {
	Path string
}

func NewJSONFile(path string) *JSONFile {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	return &JSONFile{Path: path}
}

func (j *JSONFile) Read() map[string]any {
	file, err := os.Open(j.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]any)
		}
		return make(map[string]any)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return make(map[string]any)
	}

	if len(data) == 0 {
		return make(map[string]any)
	}

	var result map[string]any
	err = json.Unmarshal(data, &result)
	if err != nil {
		return make(map[string]any)
	}

	return result
}

func (j *JSONFile) Write(data map[string]any) error {
	dir := filepath.Dir(j.Path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(j.Path, jsonData, 0644)
}
