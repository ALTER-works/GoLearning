package main

import (
	"flag"
	"fmt"
	"os"
	todo "todo-app/iternal/todo"
)

func main() {
	// Загрузка базы данных

	// Проверка, достаточно ли аргументов в вызове. Если нет, просто выводит памятку.
	if len(os.Args) < 2 {
		fmt.Println(`
 -- ToDo-app --
 Возможные команды:
 add - 
 list - 
 complete -
 delete -
 export -
 load - 
`)
		os.Exit(1)
	}

	// Обрабатка команды, в зависимости от названия и фильтров
	command := os.Args[1]

	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		descPtr := addCmd.String("desc", "", "Описание новой задачи")

		addCmd.Parse(os.Args[2:])

		t := todo.Task{Description: *descPtr, IsDone: false}
		todo.Add(t)

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		filterPtr := listCmd.String("filter", "all", "Фильтр задач")

		listCmd.Parse(os.Args[2:])

		err := todo.List(*filterPtr)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		}
	
	case "complete":
		completeCmd := flag.NewFlagSet("list", flag.ExitOnError)
		completePtr := completeCmd.String("filter", "all", "Фильтр задач")

		completeCmd.Parse(os.Args[2:])

		err := todo.List(*completePtr)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		}

	case "delete":
		deleteCmd := flag.NewFlagSet("list", flag.ExitOnError)
		deletePtr := deleteCmd.String("filter", "all", "Фильтр задач")

		deleteCmd.Parse(os.Args[2:])

		err := todo.Delete(*deletePtr)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		}

	case "export":
		exportCmd := flag.NewFlagSet("list", flag.ExitOnError)
		exportPtr := exportCmd.String("filter", "all", "Фильтр задач")

		exportCmd.Parse(os.Args[2:])

		err := todo.List(*exportPtr)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		}

	case "load":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		filterPtr := listCmd.String("filter", "all", "Фильтр задач")

		listCmd.Parse(os.Args[2:])

		err := todo.List(*filterPtr)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		}

	default:
		fmt.Printf("Неизвестная команда: %s\nИспользуйте ToDo --help для ознакомления с функционалом", command)
		os.Exit(1)
	}
}