package storage

import (
	"path/filepath"
	"testing"
	todo "todo-app/internal/todo"
)

type JsonStorage struct {
	filePath string
}

func (js JsonStorage) Load() ([]todo.Task, error){
	return LoadJSON(js.filePath)
}

func (js JsonStorage) Save(tasks []todo.Task) error{
	return SaveJSON(js.filePath, tasks)
}

func (js JsonStorage) Path() string {
	return js.filePath
}

func TestJSONStorage_Contract(t *testing.T) {
	// Создание временной папки через TempDir для json теста
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "tasks.json")

	// JSON хранилище, по которому будут проходиться тесты
	store := jsonStorageCreation(tmpFile)

	// Передача в общий тест
	verifyStorageBehavior(t, store)
}

// jsonStorageCreation создаёт экземпляр Storage из JsonStorage с указаным filePath
func jsonStorageCreation(filePath string) Storage {
	var store Storage = JsonStorage{filePath: filePath}
	return store
}