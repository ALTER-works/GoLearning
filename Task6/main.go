package main

import (
	"fmt"
	"errors"
	"os"
	"log"
	"strings"
)

func ReadProcessWrite(inputPath string, outputPath string, process func(string) (string, error)) error {
	// Чтения файла с проверкой (есть ли, можно ли прочитать, ведёт ли на папку, заблокирован ли другим процессом)
	file, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("Ошибка во время чтения: %w", err)
	}


	outputFile, err2 := process(string(file))
	if err2 != nil{
		return fmt.Errorf("Ошибка во время преобразований: %w",err2)
	}

	// Маски:
	// 0644 - чтение + запись для пользователя. Чтение для остальных.
	// 0755 - чтение + запись + запуск для пользователя. Чтение + запуск для остальных.
	// 0600 - чтение + запись только для пользователя.
	err3 := os.WriteFile(outputPath, []byte(outputFile), 0644)
	if err3 != nil {
		return fmt.Errorf("Ошибка во время записи в новый файл: %w",err3)
	}
	return nil
}

// Функция (callback) для тестирования. Как я понял, можно здесь прописать любую логику. По этому я решил с отствупами поработать.
func forTests(s string) (string, error){

	s = strings.TrimSpace(s)

	// Чисто проверка на отсутсвие данных (у readfile её нет).
	if s == ""{
		return "", errors.New("файл не содержит данных")
	}

	// Проеврка на наличие определённого текста (привет твич)
	if strings.Contains(strings.ToLower(s), "error") {
		return "", errors.New("текст содержит запрещённое слово 'error'")
	}

	// Увеличение всех i в соответствии с английской граматикой (слово "I" в анг должно быть заглавной)
	words := strings.Fields(s)
	for i := 0; i< len(words); i++ {
		if words[i] == "i"{
			words[i] = "I"
		} else if strings.HasPrefix(words[i], "i'"){
			words[i] = "I" + words[i][2:]
		} else if len(words[i]) > 1 && words[i][0] == 'i' {
			nextChar := words[i][1]
			if nextChar == ',' || nextChar == '.' || nextChar == ';' || nextChar == '?' || nextChar == '!' || nextChar == '~' || nextChar == ':' || nextChar == '*' {
				words[i] =  "I" + words[i][1:]
			}
		}
	}

	result := strings.Join(words, " ")
	
	return result, nil
}

// основной вызов всех функций
func main() {
	err := ReadProcessWrite(
		"input.txt",
		"output.txt",
		forTests,
	)

	if err != nil {
		log.Fatalf("Ошибка выполнения: %v", err)
	}
	fmt.Println("Готово!")
}