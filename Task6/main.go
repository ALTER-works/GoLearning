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
func forTests2(s string) (string, error){

	test := strings.TrimSpace(s)

	// Чисто проверка на отсутсвие данных (у readfile её нет).
	if test == ""{
		return "", errors.New("файл не содержит данных")
	}

	// Проеврка на наличие определённого текста (привет твич)
	if strings.Contains(strings.ToLower(test), "error") {
		return "", errors.New("текст содержит запрещённое слово 'error'")
	}

	
	// Увеличение всех i в соответствии с английской граматикой (слово "I" в анг должно быть заглавной)
	// Да, кстати, т.к. strings to upper делает все символы заглавными, логика этой конкретно функции будет не нужна.
	// По этому я и разделяю функции, чтобы оставить эту, но при этом сделать задание с учётом ТЗ по Upper.
	// Что касается не сохранения переносов, всё дело в логике. При разделении s на words функция также считывает переносы и табуляции как разделители.
	// В следствии чего, при разделении, они пропадают. А при дальнейшей сборке words обратно в единую строку, переносы уже не ставятся, т.к. заранее не были учтены.
	// Для учёта переносов нужно будет отдельную функцию разделения либо пакетом, либо собственными руками создавать и добавлять.
	words := strings.Fields(s)
	for i := 0; i< len(words); i++ {
		if words[i] == "i"{
			words[i] = "I"
		} else if strings.HasPrefix(words[i], "i'"){
			words[i] = "I" + words[i][1:]
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

// Второй callback, но с завёрнутым функционалом ToUpper. По тз.
func forTests(s string) (string, error){
	test := strings.TrimSpace(s)

	// Чисто проверка на отсутсвие данных (у readfile её нет).
	if test == ""{
		return "", errors.New("файл не содержит данных")
	}

	// Проеврка на наличие определённого текста (привет твич)
	if strings.Contains(strings.ToLower(test), "error") {
		return "", errors.New("текст содержит запрещённое слово 'error'")
	}

	s = strings.ToUpper(s)

	return s, nil
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
	fmt.Println("Файл сохранён в output.txt")

	err2 := ReadProcessWrite(
		"input2.txt",
		"output2.txt",
		forTests2,
	)

	if err2 != nil {
		log.Fatalf("Ошибка выполнения: %v", err2)
	}
	fmt.Println("Готово!")
	fmt.Println("Файл сохранён в output2.txt")

	// Решил проверить доп просто подстановку, как вы писали.
	// Интересно, однако, это всё работает.
	err3 := ReadProcessWrite(
		"input.txt",
		"outputExtra.txt",
		func(s string) (string, error) { return strings.ToUpper(s), nil },
	)
	if err3 != nil {
		log.Fatalf("Ошибка выполнения: %v", err3)
	}
	fmt.Println("Готово!")
	fmt.Println("Файл сохранён в outputExtra.txt")
}