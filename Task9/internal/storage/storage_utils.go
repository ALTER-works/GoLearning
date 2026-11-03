package storage

import (
	"fmt"
	"os"
	"net/http"
)

// dataLoad загружает данные из файлов для проекта через Unmarhsal из csv и json. 
func DataLoad(pathToTest string) ([]byte, error){
	dataBytes, err := os.ReadFile(pathToTest)
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