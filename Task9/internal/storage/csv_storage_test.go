package storage

import (
	"testing"
)

const pathToCSVTest string = "testdata/mock_data.csv"

func TestLoadCSV(t *testing.T){
	// Подгружаем данные

	// Сама проверяемая функция
	LoadCSV(pathToCSVTest)

	// Проверка
}

func TestSaveCSV(t *testing.T){
	// Подгружаем данные
	data, err := dataLoad(pathToCSVTest)
	if err != nil {
		t.Fatal(err)
	}

	// Сама проверяемая функция
	newPathToCSVTest := "testdata/mock_data_save.csv"
	SaveCSV(newPathToCSVTest, data)

	// Снова подгружаем данные из нового места
	newData, err := dataLoad(newPathToCSVTest)
	if err != nil {
		t.Fatal(err)
	}

	// Сверяем данные из двух файлов
	if len(data) != len(newData) {
		t.Fatalf("Ошибка: сохранённые данные не совпали! (по длинне)")
	} else {
		for i := 0; i < len(data); i++ {
			if data[i] != newData[i] {
				t.Logf("Ошибка: сохранённые данные не совпали! (На %v-ом элеметне)", i)
				t.Fatalf("%v != %v", data[i], newData[i])
			}
		}
	}
	t.Logf("Тест пройден: данные успешно сохранены по пути: %v", newPathToCSVTest)
}