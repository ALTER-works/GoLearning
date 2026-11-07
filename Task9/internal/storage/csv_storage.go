package storage

import (
	"fmt"
	"os"
	add "todo-app/internal/storage/additions"
	todo "todo-app/internal/todo"
)

func LoadCSV(path string) ([]todo.Task, error) {
	dataBytes, err := DataLoad(path)
	if err != nil {
		return nil, err
	}
	
	if len(dataBytes) == 0 {
		return []todo.Task{}, nil
	}

	var tasks []todo.Task
	err = add.CsvUnmarshal(dataBytes, &tasks)
	if err != nil {
		return nil, fmt.Errorf("не удалось анмаршалить данные из CSV: %v", err)
	}

	return tasks, nil
}

func SaveCSV(path string, tasks []todo.Task) error {
	err := dirCreation(path)
	if err != nil {
		return err
	}

	dataBytes, err := add.CsvMarshal(tasks)
	if err != nil {
		return fmt.Errorf("не удалось маршалить данные в CSV: %w", err)
	}

	err = os.WriteFile(path, dataBytes, 0664)
	if err != nil {
		return fmt.Errorf("не удалось записать CSV файл по пути %s: %w", path, err)
	}

	return nil
}