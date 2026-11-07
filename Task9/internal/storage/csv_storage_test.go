package storage

import (
	"path/filepath"
	"testing"
	todo "todo-app/internal/todo"
)

type CSVStorage struct {
	filePath string
}

func (cs CSVStorage) Load() ([]todo.Task, error){
	return LoadCSV(cs.filePath)
}

func (cs CSVStorage) Save(tasks []todo.Task) error{
	return SaveCSV(cs.filePath, tasks)
}

func (cs CSVStorage) Path() string {
	return cs.filePath
}

func TestCSVStorage_Contract(t *testing.T) {
	// Создание временной папки через TempDir для csv теста
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "tasks.csv")

	// CSV хранилище, по которому будут проходиться тесты
	store := csvStorageCreation(tmpFile)

	// Передача в общий тест
	verifyStorageBehavior(t, store)
}

// csvStorageCreation создаёт экземпляр Storage из CSVStorage с указаным filePath
func csvStorageCreation(filePath string) Storage {
	var store Storage = JsonStorage{filePath: filePath}
	return store
}