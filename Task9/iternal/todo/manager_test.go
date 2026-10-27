package todo

import (
	"encoding/json"
	"fmt"
	"testing"
	"os"
	"io"
)

const originalTestDbPath string = ""
const copyTestDbPath string = ""

const testTask string = "Testing smth new"

var testTaskStruct = Task{
	ID: 1, 
	Description: "Testing smth new x2",
	IsDone: false,
}

// TestMain - точка входа в тесты
func TestMain(m *testing.M) {

	err := cloneFile(originalTestDbPath, copyTestDbPath)
	if err != nil {
		fmt.Printf("Критическая ошибка: не удалось создать копию базы: %v\n", err)
		os.Exit(1) // прерываем выполнение, если не удалось скопировать данные
	}

	// Запуск тестов
	code := m.Run()

	// Удаление копии бд
	_ = os.Remove(copyTestDbPath)

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

func TestAdd(t *testing.T) {
	t.Log("Проверка: --Add--")
	a, err := countJsonElems(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: не удалось посчитать элементы до добавления: %v", err)
	}
	t.Logf("До добавления в файле было элементов: %d", a)

	// Add(testTaskStruct)

	b, err := countJsonElems(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: не удалось посчитать элементы после добавления: %v", err)
	}
	t.Logf("После добавления в файле стало элементов: %d", b)

	if (a + 1) != b {
		t.Fatalf("Ошибка: ожидали %d элементов, но получили %d (элемент не был добавлен)", a+1, b)
	} else {
		t.Logf("Тест пройден: элемент добавлен")
	}
}

func TestDelete(t *testing.T) {
	t.Log("Проверка: --Delete--")
	a, err := countJsonElems(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: не удалось посчитать элементы до удаления: %v", err)
	}
	t.Logf("До удаления в файле было элементов: %d", a)

	// Delete(testTaskStruct)

	b, err := countJsonElems(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: не удалось посчитать элементы после удаления: %v", err)
	}
	t.Logf("После удаления в файле стало элементов: %d", b)

	if (a - 1) != b {
		t.Fatalf("Ошибка: ожидали %d элементов, но получили %d (элемент не был добавлен)", a-1, b)
	} else {
		t.Logf("Тест пройден: элемент удалён")
	}
}

func TestList(t *testing.T) {
	t.Log("Проверка: --List--")
	s, err := List()
	if err != nil {
		t.Errorf("Ошибка при вызове List без аргументов: %v", err)
	}

	s, err = List()
	if err != nil {
		t.Errorf("Ошибка при вызове List с флагом ?: %v", err)
	}

	s, err = List()
	if err != nil {
		t.Errorf("Ошибка при вызове List с флагом ?: %v", err)
	}

	s, err = List()
	if err != nil {
		t.Errorf("Ошибка при вызове List с флагом ?: %v", err)
	}
}

func TestComplete(t *testing.T) {
	t.Log("Проверка: --Complete--")

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

	Complete(id)
}