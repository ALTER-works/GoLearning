package storage

import (
	"fmt"
	"os"
	"net/http"
	"path/filepath"
)

// dataLoad загружает данные из файлов для проекта через Unmarhsal из csv и json. Также отвечает за подгрузку данных из БД.
func DataLoad(pathToFile string) ([]byte, error){
	dataBytes, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, fmt.Errorf("Ошибка: %w", err)
	}
	
	return dataBytes, nil
}

type FileType string
// findType ищет тип файла для dataLoad
func findType(dataBytes []byte) (FileType, error)  {
	if len(dataBytes) == 0 {
		return "", fmt.Errorf("Нет данных. Количество байт в файле: %v", len(dataBytes))
	}

	limit := 512
	if len(dataBytes) < limit {
		limit = len(dataBytes)
	}
	
	mimeType := http.DetectContentType(dataBytes[:limit])
	
	return FileType(mimeType), nil
}

// dirCreation создает папку по указанному пути
func dirCreation(path string) error {
    dir := filepath.Dir(path)
    if err := os.MkdirAll(dir, os.ModePerm); err != nil {
        return fmt.Errorf("не удалось создать директорию %s: %w", dir, err)
    }

	return nil
}