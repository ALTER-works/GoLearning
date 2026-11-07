package storage

import (
	"reflect"
	"testing"
	todo "todo-app/internal/todo"
)

type Storage interface {
	Load() ([]todo.Task, error)
	Save(tasks []todo.Task) error
	Path() string
}

func verifyStorageBehavior(t *testing.T, store Storage) {
	t.Helper() // Говорит Go, что при ошибке нужно показывать строку из вызывающего теста

	// Входные данные для проверки
	input := []todo.Task{
		{ID: 1, Description: "Тест 1", IsDone: false},
		{ID: 2, Description: "Тест 2", IsDone: true},
	}

	// Проверка сохранения
	t.Log("Проверка: --Save--")
	if err := store.Save(input); err != nil {
		t.Fatalf("Ошибка сохранения: %v", err)
	}
	t.Logf("Тест пройден: данные успешно сохранены по пути %s", store.Path())

	// Проверка загрузки
	t.Log("Проверка: --Load--")
	output, err := store.Load()
	if err != nil {
		t.Fatalf("Ошибка загрузки: %v", err)
	}
	t.Logf("Тест пройден: данные успешно сохранены по пути %s", store.Path())

	// Проверка, что сохранилось ровно то, что было отправилено на вход
	if !reflect.DeepEqual(input, output) {
		t.Errorf("Данные искажены.\nОжидали: %+v\nПолучили: %+v", input, output)
	}
}