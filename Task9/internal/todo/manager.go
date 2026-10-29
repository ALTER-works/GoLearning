package todo

import "fmt"

// Add - функция, добавляющая объект task в tasks. Возвращает изменённый tasks соотв-но.
func Add(tasks []Task, desc string) ([]Task, error) {
	newTask := Task{
		ID: randomIdGenerator(len(tasks)),
		Description: desc,
		IsDone: false,
	}

	tasks = append(tasks, newTask)
	return tasks, nil
}

// List - функция, выводящая все данные из базы (все задачи). Возвращает тот же tasks.
func List(tasks []Task, filter string) ([]Task, error) {
	if (len(tasks)) == 0 {
		return tasks, fmt.Errorf("Внимание: переданный массив не имеет данных")
	}

	if filter == "" || filter == "all"{
		return tasks, nil
	}

	var filteredTasks []Task

	switch  filter{
	case "done":
		for _, task := range tasks {
			if task.IsDone == true {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case "pending":
		for _, task := range tasks {
			if task.IsDone == false {
				filteredTasks = append(filteredTasks, task)
			}
		}
	default:
		return tasks, fmt.Errorf("Ошибка: неизвестны фильтр %v", filter)
	}
	return tasks, nil
}

// Complete - функция, меняющая значение IsDone объекта Task на противоположное. Возвращает изменённый tasks соотв-но.
func Complete(tasks []Task, id int) ([]Task, error) {
	
	foundId := findTaskIndex(tasks, id)
	if foundId == -1 {
		return tasks, fmt.Errorf("Ошибка: задача с ID %v не найдена", id)
	}

	tasks[foundId].IsDone = !tasks[foundId].IsDone
	return tasks, nil
}

// Delete - функция, удаляющая объект task из массива tasks. Возвращает изменённый tasks соотв-но.
func Delete(tasks []Task, id int) ([]Task, error) {
	foundId := findTaskIndex(tasks, id)
	if foundId == -1 {
		return tasks, fmt.Errorf("Ошибка: задача с ID %v не найдена", id)
	}

	tasks = append(tasks[:foundId], tasks[foundId + 1:]...)

	return tasks, nil
}

// findTaskIndex - Вспомогательная функция для поиска индекса в Complete и Delete
func findTaskIndex(tasks []Task, id int) int {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			return i
		}
	}
	return -1
}
