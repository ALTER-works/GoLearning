package main

import (
	"fmt";
	"errors";
	"os";
	"log";
	"strings";
)

func ReadProcessWrite(inputPath string, outputPath string, process func(string) (string, error)) error {
	textFromFile, err := os.ReadFile(inputPath)
	

	if err != nil {
		return err
	}
	process(textFromFile)
	return nil
}

func main() {
	err := ReadProcessWrite(
		"inputText",
		"outputText",

	)

	if err != nil {
		log.Fatal("Ошибка выполнения: %v", err)
	}
	fmt.Println("Готово!")
}