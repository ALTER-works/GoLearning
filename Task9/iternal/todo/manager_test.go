package todo

import (
	"encoding/json"
	"fmt"
	"testing"
	"os"
	"io"
)

const originalTestDbPath string = "../testdata/test_tasks.json"
const copyTestDbPath string = "../testdata/test_tasks_copy.json"

// TestMain - точка входа в тесты
func TestMain(m *testing.M) {

	// Клонирование тестовых данных
	err := cloneFile(originalTestDbPath, copyTestDbPath)
	if err != nil {
		fmt.Printf("Критическая ошибка: не удалось создать копию базы: %v\n", err)
		os.Exit(1)
	}

	// Запуск тестов
	code := m.Run()

	// Удаление копии бд
	err = os.Remove(copyTestDbPath)
	if err != nil {
		fmt.Printf("Критическая ошибка: не удалось удалить копию базы: %v\n", err)
		os.Exit(1)
	}

	// Выход
	os.Exit(code)
}

// cloneFile - клонирует копию тестовой бд, чтобы при вызове тестов основную бд не менять
func cloneFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

// countJsonElems - считает кол-во элементов в json файле, а потом выводит итоговое кол-во
func countJsonElems(copyTestDbPath string) (int, error) {
	data, err := os.ReadFile(copyTestDbPath)
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения файла %s: %w", copyTestDbPath, err)
	}

	var items []json.RawMessage
	
	err = json.Unmarshal(data, &items)
	if err != nil {
		return 0, fmt.Errorf("ошибка десериализации JSON: %w", err)
	}
	
	return len(items), nil
}

// TestAdd - проверка функции Add из manage.go
func TestAdd(t *testing.T) {
	t.Log("Проверка: --Add--")
	a, err := countJsonElems(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: не удалось посчитать элементы до добавления: %v", err)
	}
	t.Logf("До добавления в файле было элементов: %d", a)

	// Запуск функции
	err = Add(id)
	if err != nil {
		t.Logf("Ошибка: %v", err)
	}

	b, err := countJsonElems(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: не удалось посчитать элементы после добавления: %v", err)
	}
	t.Logf("После добавления в файле стало элементов: %d", b)

	// Проверка на то, увеличилось ли кол-во элементов
	if (a + 1) != b {
		t.Fatalf("Ошибка: ожидали %d элементов, но получили %d", a+1, b)
	} else {
		t.Logf("Тест пройден: элемент добавлен")
	}
}

// TestDelete - проверка функции Delete из manage.go
func TestDelete(t *testing.T) {
	t.Log("Проверка: --Delete--")
	a, err := countJsonElems(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: не удалось посчитать элементы до удаления: %v", err)
	}
	t.Logf("До удаления в файле было элементов: %d", a)

	// Запуск функции
	err = Delete(id)
	if err != nil {
		t.Logf("Ошибка: %v", err)
	}		

	b, err := countJsonElems(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: не удалось посчитать элементы после удаления: %v", err)
	}
	t.Logf("После удаления в файле стало элементов: %d", b)

	// Проверка на то, уменьшилось ли кол-во элементов
	if (a - 1) != b {
		t.Fatalf("Ошибка: ожидали %d элементов, но получили %d (элемент не был добавлен)", a-1, b)
	} else {
		t.Logf("Тест пройден: элемент удалён")
	}
}

// TestList - проверка функции List из manage.go
func TestList(t *testing.T) {
	t.Log("Проверка: --List--")

	s, err := List(copyTestDbPath)
	if err != nil {
		t.Errorf("Ошибка при вызове List без аргументов: %v", err)
	}


}

// TestComplete - проверка функции Complete из manage.go
func TestComplete(t *testing.T) {
	t.Log("Проверка: --Complete--")

	// Создание массива данных из чтения байтов копии тестового json файла
	data, err := os.ReadFile(copyTestDbPath)
	if err != nil {
		t.Fatalf("ошибка чтения файла %s: %v", copyTestDbPath, err)
	}

	var items []Task
	err = json.Unmarshal(data, &items)
	if err != nil {
		t.Fatalf("ошибка десериализации JSON: %v", err)
	}

	if len(items) == 0 {
		t.Fatalf("ошибка: тестовая база пуста")
	}

	// Тут находим последний элемент этого массива, 
	// по чьему айди мы будем проводить изменения с помощью функции Complete в тесте
	lastIndex := len(items) - 1
	lastItem := items[lastIndex]
	
	id := lastItem.ID
	initialIsDone := lastItem.IsDone

	t.Logf("Последний элемент: ID=%d, Статус IsDone=%t", id, initialIsDone)

	// Запуск функции
	err = Complete(id)
	if err != nil {
		t.Logf("Ошибка: %v", err)
	}

	// Создание нового массива из изменённых данных
	newData, err := os.ReadFile(copyTestDbPath)
	if err != nil {
		t.Fatalf("ошибка перечитывания файла после Complete: %v", err)
	}

	var updatedItems []Task
	err = json.Unmarshal(newData, &updatedItems)
	if err != nil {
		t.Fatalf("ошибка повторной десериализации JSON: %v", err)
	}

	// Тут аналогично находим последний элемент + выводим
	newIsDone := updatedItems[lastIndex].IsDone
	t.Logf("Статус после Complete: IsDone=%t", newIsDone)

	// Сравниваем результаты - если после complete ничего не поменялось по айди, выдаём ошибку
	if initialIsDone == newIsDone {
		t.Errorf("Ошибка: после вызова Complete статус IsDone для ID %v не изменился!", id)
	}else {
		t.Logf("Тест пройден: значение IsDone у Task (%v) из DB (%v) успешно изменено с %v на %v", lastItem, copyTestDbPath, initialIsDone, newIsDone)
	}
}