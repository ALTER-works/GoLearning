package storage

import (
	"fmt"
	"os"
	todo "todo-app/internal/todo"
)

func LoadCSV(path string) ([]todo.Task, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Не удалось загрузить данные из пути: %v.\nОшибка: %v", path, err)
	}


	return nil, nil
}

func SaveCSV(path string, tasks []todo.Task) error {
	return nil
}