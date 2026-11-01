package storage

import (
	"fmt"
	todo "todo-app/internal/todo"
)

func LoadCSV(path string) ([]todo.Task, error) {
	tasks, err := DataLoad(path)
	if err != nil {
		return nil, fmt.Errorf("Ошибка загрузки файла: %v", err)
	}

	return nil, nil
}

func SaveCSV(path string, tasks []todo.Task) error {
	return nil
}