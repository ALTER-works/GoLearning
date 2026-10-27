package todo

// Add - функция, добавляющая элемент (task) в базу данных (файл tasks.json)
func Add(tasks []Task, desc string) []Task {
	return nil
}

// List - функция, выводящая все данные из базы (все задачи). 
// Имеет несколько флагов:  
func List(tasks []Task, filter string) []Task {
	return nil
}

// Complete - функция, меняющая значение IsDone объекта Task на противоположное
func Complete(tasks []Task, id int) ([]Task, error) {
	return nil, nil
}

// Delete - функция, удаляющая элемент (task) из базы данных (файл tasks.json) 
func Delete(tasks []Task, id int) ([]Task, error) {
	return nil, nil
}