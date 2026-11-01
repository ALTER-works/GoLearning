package storage

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	todo "todo-app/internal/todo"
)

type FileType string

const (
	TypeJSON    FileType = "json"
	TypeCSV     FileType = "csv"
	TypeUnknown FileType = "unknown"
)

func dataLoad(pathToTest string) ([]todo.Task, error){
	dataBytes, err := os.ReadFile(pathToTest)
	if err != nil {
		return nil, fmt.Errorf("Ошибка: %v", err)
	}

	kind, err := findType(dataBytes)
	if err != nil {
		return nil, fmt.Errorf("Ошибка: %v", err)
	}

	var data []todo.Task
	switch kind {
	case "csv":
		err = csvUnmarshal(dataBytes, &data)
		if err != nil {
			return nil, fmt.Errorf("Ошибка: %v", err)
		}
	case "json":
		err = json.Unmarshal(dataBytes, &data)
		if err != nil {
			return nil, fmt.Errorf("Ошибка: %v", err)
		}
	default:
		return nil, fmt.Errorf("Ошибка: неподходящий тип файла")
	}

	return data, nil
}

func findType(dataBytes []byte) (FileType, error){
	trimmed := bytes.TrimSpace(dataBytes)
	if len(trimmed) == 0 {
		return TypeUnknown, fmt.Errorf("Ошибка: неподходящий тип файла")
	}


	return "", nil
}

// Решил написать, т.к. аналога Unmarshal базового для csv в Go нет, 
// а использовать модуль из стороннего проекта не захотелось.

// csvUnmarshal - аналог json.unmarshal, только для csv. Реализован через рефлексию.
func csvUnmarshal(dataBytes []byte, v any) error {
	// Проверка на то, пустой ли файл или нет, как в unmarshal json
	trimmed := bytes.TrimSpace(dataBytes)
	if len(trimmed) == 0 {
		return fmt.Errorf("длинна файла равна %v", len(trimmed))
	}

	// Проверка на то, является v - указателем на слайс, т.к. нам нужно именно изначальные данные менять
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return fmt.Errorf("csv: Unmarshal ожидает ненулевой указатель, получено %T", v)
	}

	// Получаем наш слайс в обёртке интерфейса рефлексии
	sliceVal := rv.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return fmt.Errorf("csv: Unmarshal ожидает указатель на слайс, получен указатель на %s", sliceVal.Kind())
	}

	// Получаем схему значения, лежащего в слайсе
	elemType := sliceVal.Type().Elem()

	// Считывание данных из файла через базовые модули csv и bytes
	reader := csv.NewReader(bytes.NewReader(dataBytes))
	// Себе на будущее: 
	// в го чтение и обработка байтов (парсинг) разделены на два слоя - io.reader (чтение данных, интерфейс) и decoder (обработку)
	// Эти слои состоят из разных функций из разных модулей, но в совокупности решают одну задачу - прочитать информацию и обработать её (привести к какому-то типу, к какой-то форме)
	// У обработчика (тут csv.NewReader) нет своего интерфейса, т.к. для разных типов файлов нужна разная обработка.
	// Reader csv не умеет работать с другими кодировкам, он работает только с utf-8 и соотв. ASCII 
	
	var header []string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("csv read: %v", err)
		}

		if header == nil {
			header = record
			continue // Пропускаем строку заголовков
		}
		
		// Создаём новый элемент на основе типа через рефлексию
		newElem := reflect.New(elemType).Elem()
		for i := 0; i < len(header); i++ {
			field := newElem.Field(i)
			if !field.CanSet() {
				// fmt.Printf("Внимание: поле '%v' не экспортируемое\n", header[i])
				continue
			}

			fieldType := field.Type()
			rawValue := record[i]

			switch fieldType.Kind() {
			case reflect.String:
				field.SetString(rawValue)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				val, err := strconv.ParseInt(rawValue, 10, 64)
				if err != nil {
					fmt.Printf("Ошибка конвертации поля %s: %v\n", header[i], err)
					continue
				}
				field.SetInt(val)
			case reflect.Float32, reflect.Float64:
				val, err := strconv.ParseFloat(rawValue, 64)
				if err != nil {
					fmt.Printf("Ошибка конвертации поля %s: %v\n", header[i], err)
					continue
				}
				field.SetFloat(val)
			case reflect.Bool:
				val, err := strconv.ParseBool(rawValue)
				if err != nil {
					fmt.Printf("Ошибка конвертации поля %s: %v\n", header[i], err)
					continue
				}
				field.SetBool(val)
			}
		}

		sliceVal.Set(reflect.Append(sliceVal, newElem))
	}
	
	return nil
}

// type Task struct {
// 	ID int
// 	Description string
// 	IsDone bool
// }