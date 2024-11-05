package filedto

import (
	"fmt"
	"goph-keeper/internal/helpers"
	"goph-keeper/internal/logger"
	"io"
	"mime/multipart"
	"net/http"
)

type SaveFileDTO struct {
	NameFile string `json:"name_file"` // Описание в байтовом формате
	FileData []byte `json:"file_data"` // Содержимое файла
	UserID   int    `json:"user_id"`
}

func SaveFileDTOFromHTTP(r *http.Request) (SaveFileDTO, error) {
	// Получаем файл из формы
	file, _, err := r.FormFile("file")
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error retrieving file.proto from form-data: %w", err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			logger.Log.Errorf("error closing file.proto: %v", err)
		}
	}(file)

	// Читаем содержимое файла
	fileData, err := io.ReadAll(file)
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error reading file.proto: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error retrieving userID from context: %w", err)
	}

	// Читаем JSON-данные из формы
	nameFile := r.FormValue("name_file")

	// Формируем DTO с данными
	saveFileDTO := SaveFileDTO{
		UserID:   userID,
		NameFile: nameFile,
		FileData: fileData,
	}

	return saveFileDTO, nil
}
