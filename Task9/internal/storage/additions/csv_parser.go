package storage

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
	"strconv"
)
// Решил написать, т.к. аналога Unmarshal базового для csv в Go нет, 
// а использовать модуль из стороннего проекта не захотелось.

// CsvUnmarshal переводит байты файла CSV сразу в необходимый тип данных по возможности. Аналог json.unmarshal, только для csv. Реализован через рефлексию.
func CsvUnmarshal(dataBytes []byte, v any) error {
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
	
	// Технические переменные для цикла
	var header []string
	fastSlice := reflect.MakeSlice(sliceVal.Type(), 0, 0)
	flagToStop := true

	for flagToStop {
		record, err := reader.Read()
		if err == io.EOF {
			flagToStop = false
			continue
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

		fastSlice = reflect.Append(fastSlice, newElem)
	}

	sliceVal.Set(fastSlice)

	return nil
}

// CsvMarshal переводит данные из опреденённого типа в байты для записи в файл. Аналог json.marshal и обратная копия unmarshal. Тоже рефлексия.
// Нужно будет подкорректировать в будущем некоторые моменты, т.к. это копия через МЛ
func CsvMarshal(v any) ([]byte, error) {
	// Проверка на то, что нам передали не пустые данные (если слайс пустой, возвращаем пустой массив байт)
	if v == nil {
		return nil, fmt.Errorf("csv: Marshal ожидает данные, получено nil")
	}

	// Проверка на то, является ли v слайсом или указателем на слайс, т.к. маршалить можно и то, и другое
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil, fmt.Errorf("csv: Marshal ожидает ненулевой указатель, получено %T", v)
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("csv: Marshal ожидает слайс или указатель на слайс, получено %s", rv.Kind())
	}

	// Получаем схему значения, лежащего в слайсе
	elemType := rv.Type().Elem()
	if elemType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("csv: Marshal ожидает слайс структур, получен слайс %s", elemType.Kind())
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Технические переменные для цикла
	var header []string
	length := rv.Len()

	// Собираем заголовки на основе экспортируемых полей структуры
	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		if field.PkgPath == "" {
			header = append(header, field.Name)
		}
	}

	// Записываем строку заголовков в буфер
	if err := writer.Write(header); err != nil {
		return nil, fmt.Errorf("csv write header: %v", err)
	}

	for i := 0; i < length; i++ {
		structVal := rv.Index(i)
		var record []string

		for j := 0; j < elemType.NumField(); j++ {
			fieldInfo := elemType.Field(j)
			if fieldInfo.PkgPath != "" {
				continue // Пропускаем неэкспортируемые поля
			}

			field := structVal.Field(j)
			var rawValue string

			switch field.Kind() {
			case reflect.String:
				rawValue = field.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				rawValue = strconv.FormatInt(field.Int(), 10)
			case reflect.Float32, reflect.Float64:
				rawValue = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.Bool:
				rawValue = strconv.FormatBool(field.Bool())
			default:
				// Если тип сложный, приводим к строке через базовый спринт
				rawValue = fmt.Sprintf("%v", field.Interface())
			}

			record = append(record, rawValue)
		}

		if err := writer.Write(record); err != nil {
			return nil, fmt.Errorf("csv write record: %v", err)
		}
	}

	// Сбрасываем оставшиеся в буфере csv.Writer данные в наш bytes.Buffer
	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, fmt.Errorf("csv flush: %v", err)
	}

	return buf.Bytes(), nil
}