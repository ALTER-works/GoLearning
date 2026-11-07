package todo

import (
	"encoding/json"
	"fmt"
	"testing"
	"os"
	"io"
)

const originalTestDbPath string = "testdata/test_tasks.json"
const copyTestDbPath string = "testdata/test_tasks_copy.json"
const testDbDone string = "testdata/test_done_tasks.json"
const testDbPending string = "testdata/test_pending_tasks.json"

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

// tasksArrCreation - создаёт tasks []Task из копии тестовых данных
func tasksArrCreation(path string) ([]Task, error) {
	// Создание массива данных из чтения байтов копии тестового json файла
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %s: %v", path, err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("ошибка десериализации JSON: %v", err)
	}

	if len(tasks) == 0 {
		return nil, fmt.Errorf("ошибка: тестовая база пуста")
	}

	return tasks, nil
}

// TestList - проверка функции List из manage.go
func TestList(t *testing.T) {
	t.Log("Проверка: --List--")

	// Создание массива данных через tasksArrCreation
	tasks, err := tasksArrCreation(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// Проверка без фильтра
	filter1 := ""

	newTasks, err := List(tasks, filter1)
	if err != nil {
		t.Fatalf("Ошибка при вызове List без фильтра: %v", err)
	}

	if len(newTasks) != len(tasks) { 
		t.Logf("Длинна сверяемых данных: %v, длинна новых: %v", len(tasks), len(newTasks))
		t.Fatalf("Ошибка: возвращённые данные при обработке теста без фильтра не совпали (по длинне)")
	} else {
		for i := 0; i < len(newTasks); i++ {
			if tasks[i] != newTasks[i] {
				t.Fatalf("Ошибка: возвращённые данные при обработке теста без фильтра не совпали (по значениям: %v != %v)", tasks[i], newTasks[i])
			}
		}
	}
	t.Logf("Тест пройден: возвращённые данные при обработке теста без фильтра совпали")

	// Проверка с фильтром --all
	filter2 := "all"

	newTasks, err = List(tasks, filter2)
	if err != nil {
		t.Fatalf("Ошибка при вызове List с фильтром --all: %v", err)
	}

	if len(newTasks) != len(tasks) { 
		t.Logf("Длинна сверяемых данных: %v, длинна новых: %v", len(tasks), len(newTasks))
		t.Fatalf("Ошибка: возвращённые данные при обработке теста с фильтром --all не совпали (по длинне)")
	} else {
		for i := 0; i < len(newTasks); i++ {
			if tasks[i] != newTasks[i] {
				t.Fatalf("Ошибка: возвращённые данные при обработке теста с фильтром --all не совпали (по значениям: %v != %v)", tasks[i], newTasks[i])
			}
		}
	}
	t.Logf("Тест пройден: возвращённые данные при обработке теста с фильтром --all совпали")

	// Проверка с фильтром --done
	doneTasks, err := tasksArrCreation(testDbDone)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}
	filter3 := "done"

	newTasks, err = List(tasks, filter3)
	if err != nil {
		t.Fatalf("Ошибка при вызове List с фильтром --done: %v", err)
	}

	if len(newTasks) != len(doneTasks) { 
		t.Logf("Длинна сверяемых данных: %v, длинна новых: %v", len(doneTasks), len(newTasks))
		t.Fatalf("Ошибка: возвращённые данные при обработке теста с фильтром --done не совпали (по длинне)")
	} else {
		for i := 0; i < len(newTasks); i++ {
			if doneTasks[i] != newTasks[i] {
				t.Fatalf("Ошибка: возвращённые данные при обработке теста с фильтром --done не совпали (по значениям: %v != %v)", doneTasks[i], newTasks[i])
			}
		}
	}
	t.Logf("Тест пройден: возвращённые данные при обработке теста с фильтром --done совпали")

	// Проверка с фильтром --pending
	pendingTasks, err := tasksArrCreation(testDbPending)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}
	filter4 := "pending"

	newTasks, err = List(tasks, filter4)
	if err != nil {
		t.Fatalf("Ошибка при вызове List с фильтром --pending: %v", err)
	}

	if len(newTasks) != len(pendingTasks) { 
		t.Logf("Длинна сверяемых данных: %v, длинна новых: %v", len(pendingTasks), len(newTasks))
		t.Fatalf("Ошибка: возвращённые данные при обработке теста с фильтром --pending не совпали (по длинне)")
	} else {
		for i := 0; i < len(newTasks); i++ {
			if pendingTasks[i] != newTasks[i] {
				t.Fatalf("Ошибка: возвращённые данные при обработке теста с фильтром --pending не совпали (по значениям: %v != %v)", doneTasks[i], newTasks[i])
			}
		}
	}
	t.Logf("Тест пройден: возвращённые данные при обработке теста с фильтром --pending совпали")
}

// TestAdd - проверка функции Add из manage.go
func TestAdd(t *testing.T) {
	t.Log("Проверка: --Add--")

	// Создание массива данных через tasksArrCreation
	tasks, err := tasksArrCreation(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// Описание для нового task
	desc := "Тестирование добавления"

	// Запуск функции
	newTasks, err := Add(tasks, desc)
	if err != nil {
		t.Logf("Ошибка: %v", err)
	}		
	
	oldTaksLen := len(tasks)
	if len(newTasks) != (oldTaksLen+1) {
		t.Errorf("Ошибка: длинна итогового массива значений не увеличилась ровно на 1")
	} else {
		dFlag := false
		index := 0
		for (dFlag == false) && (oldTaksLen > index) {
			if newTasks[index] != tasks[index] {
				dFlag = true
			} else {
				index++
			}
		}
		for i := index; i < oldTaksLen; i++ {
			if tasks[i] != newTasks[i+1] {
				t.Errorf("Ошибка: данные после смещения не совпали на индексе %d", i)
			}
		}
	}
	t.Logf("Тест пройден: элемент добавлен")
}

// TestDelete - проверка функции Delete из manage.go
func TestDelete(t *testing.T) {
	t.Log("Проверка: --Delete--")

	// Создание массива данных через tasksArrCreation
	tasks, err := tasksArrCreation(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// Аналогично функии Complete, но берём для теста элемент из середины +-
	lastIndex := len(tasks)/2
	lastTask := tasks[lastIndex]
	id := lastTask.ID

	// Запуск функции
	newTaks, err := Delete(tasks, id)
	if err != nil {
		t.Logf("Ошибка: %v", err)
	}		
	
	newTaksLen := len(newTaks)
	if newTaksLen != (len(tasks)-1) {
		t.Errorf("Ошибка: длинна итогового массива значений не сократилась ровно на 1")
	} else {
		dFlag := false
		index := 0
		for (dFlag == false) && (newTaksLen > index) {
			if newTaks[index] != tasks[index] {
				dFlag = true
			} else {
				index++
			}
		}
		for i := index; i < newTaksLen; i++ {
			if newTaks[i] != tasks[i+1] {
				t.Errorf("Ошибка: данные после смещения не совпали на индексе %d", i)
			}
		}
	}
	t.Logf("Тест пройден: элемент удалён")
}

// TestComplete - проверка функции Complete из manage.go
func TestComplete(t *testing.T) {
	t.Log("Проверка: --Complete--")

	// Создание массива данных через tasksArrCreation
	tasks, err := tasksArrCreation(copyTestDbPath)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	// Тут находим последний элемент этого массива, 
	// по чьему айди мы будем проводить изменения с помощью функции Complete в тесте
	lastIndex := len(tasks) - 1
	lastTask := tasks[lastIndex]
	id := lastTask.ID

	// Первый запуск Complete
	t.Logf("Тестируем элемент: ID=%d, Исходный статус IsDone=%t", id, lastTask.IsDone)
	firstRunTasks, err := Complete(tasks, id)
	if err != nil {
		t.Fatalf("Ошибка при первом вызове Complete: %v", err)
	}

	// Проверка изменения статуса на true
	if !firstRunTasks[lastIndex].IsDone {
		t.Fatalf("Ошибка: после первого вызова Complete статус IsDone для ID %d остался false!", id)
	}

	// Проверка на идемпотентность (Второй запуск Complete)
	t.Log("Проверка идемпотентности: вызываем Complete повторно для того же ID")
	secondRunTasks, err := Complete(firstRunTasks, id)
	if err != nil {
		t.Fatalf("Ошибка при повторном вызове Complete: %v", err)
	}

	// Проверяем, что повторный вызов ничего не испортил и статус остался true
	if !secondRunTasks[lastIndex].IsDone {
		t.Fatalf("Ошибка идемпотентности: после повторного вызова Complete статус сбросился в false!")
	}

	t.Logf("Тест пройден: функция успешно устанавливает IsDone=true и сохраняет его при повторных вызовах")
}