package storage

import (
	"fmt"
	"os"
	"encoding/json"
	todo "todo-app/internal/todo"
)

func LoadJSON(path string) ([]todo.Task, error) {
	dataBytes, err := DataLoad(path)
	if err != nil {
		return nil, err
	}
	
	var tasks []todo.Task
	err = json.Unmarshal(dataBytes, &tasks)
	if err != nil {
		return nil, fmt.Errorf("не удалось анмаршалить данные из JSON: %v", err)
	}
	return tasks, nil
}

func SaveJSON(path string, tasks []todo.Task) error {
	dataBytes, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("не удалось маршалить данные в JSON: %w", err)
	}

	err = os.WriteFile(path, dataBytes, 0664)
	if err != nil {
		return fmt.Errorf("не удалось записать JSON файл по пути %s: %w", path, err)
	}

	return nil
}