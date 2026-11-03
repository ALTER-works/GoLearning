package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"todo-app/internal/storage" // Путь к вашему пакету storage, проверьте корректность импорта
	todo "todo-app/internal/todo"
)

const pathToData = "data/tasks.json"

func main() {
	// Загрузка базы данных из JSON файла при старте
	tasks, err := storage.LoadJSON(pathToData)
	if err != nil {
		// Если файла ещё нет, инициализируем пустой слайс, иначе выходим с ошибкой
		if os.IsNotExist(err) {
			tasks = []todo.Task{}
		} else {
			fmt.Printf("Ошибка загрузки базы данных: %v\n", err)
			os.Exit(1)
		}
	}

	// Проверка, достаточно ли аргументов в вызове. Если нет, просто выводит памятку.
	if len(os.Args) < 2 {
		fmt.Println(`
 -- ToDo-app --
 Возможные команды:
 add      - Добавить новую задачу (--desc "описание")
 list     - Показать список задач (--filter "all|done|pending")
 complete - Отметить задачу как выполненную (--id int)
 delete   - Удалить задачу (--id int)
 export   - Выгрузить базу данных в CSV файл (--format "csv|json", --out "folder")
 load     - Загрузить базу данных из CSV файла (--file "tasksImportFrom.csv|.json")`)
		os.Exit(1)
	}

	// Обработка команды, в зависимости от названия и фильтров
	command := os.Args[1]

	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		descPtr := addCmd.String("desc", "", "Описание новой задачи")

		addCmd.Parse(os.Args[2:])

		if *descPtr == "" {
			fmt.Println("Ошибка: Описание задачи не может быть пустым. Используйте флаг --desc")
			os.Exit(1)
		}

		// Вызываем менеджер и обновляем наш слайс
		tasks, err = todo.Add(tasks, *descPtr)
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
			os.Exit(1)
		}

		// Загрузка данных в БД (если данные поменялись)
		errDB := storage.SaveJSON(pathToData, tasks)
		if errDB != nil {
			fmt.Printf("Критическая ошибка сохранения базы данных: %v\n", errDB)
			os.Exit(1)
		}

		fmt.Println("Задача успешно добавлена!")

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		filterPtr := listCmd.String("filter", "all", "Фильтр задач (all, done, pending)")

		listCmd.Parse(os.Args[2:])

		// Получаем отфильтрованный список задач
		filteredTasks, err := todo.List(tasks, *filterPtr)
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
			os.Exit(1)
		}
		
		// Выводим результат пользователю
		if len(filteredTasks) == 0 {
			fmt.Println("Список задач пуст.")
		} else {
			fmt.Printf("\n--- Список задач (Фильтр: %s) ---\n", *filterPtr)
			for _, t := range filteredTasks {
				status := "[ ]"
				if t.IsDone {
					status = "[X]"
				}
				fmt.Printf("%d. %s %s\n", t.ID, status, t.Description)
			}
		}

	case "complete":
		completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
		idPtr := completeCmd.Int("id", 0, "ID задачи для отметки выполнения")

		completeCmd.Parse(os.Args[2:])

		if *idPtr == 0 {
			fmt.Println("Ошибка: Укажите корректный ID задачи через флаг --id")
			os.Exit(1)
		}

		// Вызываем функцию compelte
		updatedTasks, err := todo.Complete(tasks, *idPtr)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			os.Exit(1)
		}

		// Перезаписываем БД и меняем флаг
		tasks = updatedTasks

		// Загрузка данных в БД (если данные поменялись)
		errDB := storage.SaveJSON(pathToData, tasks)
		if errDB != nil {
			fmt.Printf("Критическая ошибка сохранения базы данных: %v\n", errDB)
			os.Exit(1)
		}
		
		fmt.Printf("Задача с ID-%d успешно выполнена!\n", *idPtr)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		idPtr := deleteCmd.Int("id", 0, "ID задачи для удаления")

		deleteCmd.Parse(os.Args[2:])

		if *idPtr == 0 {
			fmt.Println("Ошибка: Укажите корректный ID задачи через флаг --id")
			os.Exit(1)
		}

		// Вызываем функцию delete
		updatedTasks, err := todo.Delete(tasks, *idPtr)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			os.Exit(1)
		}

		// Перезаписываем БД и меняем флаг
		tasks = updatedTasks

		// Загрузка данных в БД (если данные поменялись)
		errDB := storage.SaveJSON(pathToData, tasks)
		if errDB != nil {
			fmt.Printf("Критическая ошибка сохранения базы данных: %v\n", errDB)
			os.Exit(1)
		}
		
		fmt.Printf("Задача №%d успешно удалена.\n", *idPtr)

	case "export":
		exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
		formatPtr := exportCmd.String("format", "csv", "Формат выгрузки (csv|json)")
		outFolderPtr := exportCmd.String("out", "", "Папка для выгрузки файла")

		exportCmd.Parse(os.Args[2:])

		if len(*outFolderPtr) == 0 {
			fmt.Println("Использование: todo-app.exe export --format=\"csv|json\" --out=\"folder\"")
			os.Exit(1)
		}

		format := strings.ToLower(*formatPtr)
		if format != "csv" && format != "json" {
			fmt.Printf("Ошибка: неподдерживаемый формат '%s'. Используйте 'csv' или 'json'\n", format)
			os.Exit(1)
		}

		var fullPath string
		var dir string

		// Проверяем, указал ли пользователь расширение файла в конце пути
		ext := strings.ToLower(filepath.Ext(*outFolderPtr))
		
		if ext != "" {
            // Если расширение есть, оно ОБЯЗАНО быть либо .json, либо .csv
            if ext == ".json" || ext == ".csv" {
                fullPath = *outFolderPtr
                dir = filepath.Dir(fullPath)
                format = ext[1:] // автоматически переключаем формат под расширение (json или csv)
            } else {
                // Если пользователь пытается подсунуть .cpp, .txt или любой другой неподдерживаемый формат, возвращаем ошибку
                fmt.Printf("Ошибка: неподдерживаемое расширение файла '%s'. Используйте только .json или .csv\n", ext)
                os.Exit(1)
            }
        } else {
			// Если передана просто папка, добавляем дефолтное имя (tasksExported.<format>)
			dir = *outFolderPtr
			fileName := fmt.Sprintf("tasksExported.%s", format)
			fullPath = filepath.Join(dir, fileName)
		}

		// Создаем папку по указанному в вызове пути, если таковой ещё нет
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Ошибка при создании папки '%s': %v\n", dir, err)
			os.Exit(1)
		}

		// Вызываем нужный метод сохранения в зависимости от формата
		switch format {
		case "json":
			err := storage.SaveJSON(fullPath, tasks)
			if err != nil {
				fmt.Printf("Ошибка экспорта в JSON: %v\n", err)
				os.Exit(1)
			}
		case "csv":
			err := storage.SaveCSV(fullPath, tasks)
			if err != nil {
				fmt.Printf("Ошибка экспорта в CSV: %v\n", err)
				os.Exit(1)
			}
		}
		
		fmt.Printf("Данные успешно экспортированы в файл: %s\n", fullPath)

	case "load":
		loadCmd := flag.NewFlagSet("load", flag.ExitOnError)
		filePathPtr := loadCmd.String("file", "", "Путь к импортируемому CSV или JSON файлу")

		loadCmd.Parse(os.Args[2:])

		if len(*filePathPtr) == 0{
			fmt.Println("Использование: todo-app.exe load --file=<путь_импорта_с_расширением_файла>")
			os.Exit(1)
		}

		var loadedTasks []todo.Task
		ext := strings.ToLower(filepath.Ext(*filePathPtr))

		switch ext {
		case ".csv":
			// Загружаем задачи из CSV
			loadedTasks, err = storage.LoadCSV(*filePathPtr)
			if err != nil {
				fmt.Printf("Ошибка импорта из CSV: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Успешно загружено %d задач из CSV файла!\n", len(loadedTasks))

		case ".json":
			// Загружаем задачи из JSON
			loadedTasks, err = storage.LoadJSON(*filePathPtr) 
			if err != nil {
				fmt.Printf("Ошибка импорта из JSON: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Успешно загружено %d задач из JSON файла!\n", len(loadedTasks))

		default:
			fmt.Printf("Ошибка: Неподдерживаемый формат файла '%s'. Поддерживаются только .csv и .json\n", ext)
			os.Exit(1)
		}

		// Перезаписываем БД и меняем флаг
		tasks = loadedTasks

		// Загрузка данных в БД (если данные поменялись)
		errDB := storage.SaveJSON(pathToData, tasks)
		if errDB != nil {
			fmt.Printf("Критическая ошибка сохранения базы данных: %v\n", errDB)
			os.Exit(1)
		}

	default:
		fmt.Printf("Неизвестная команда: %s\nЗапустите приложение без аргументов для вывода справки.\n", command)
		os.Exit(1)
	}
}
