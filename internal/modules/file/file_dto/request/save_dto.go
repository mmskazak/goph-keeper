package request

import (
	"fmt"
	"gophKeeper/internal/helpers"
	"io"
	"net/http"
)

type SaveFileDTO struct {
	FileName string `json:"file_name"`
	FileData []byte `json:"file_data"` // Содержимое файла в байтовом формате
	UserID   int    `json:"user_id"`
}

func SaveFileDTOFromHTTP(r *http.Request) (SaveFileDTO, error) {
	// Ограничиваем размер загружаемого файла
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error parsing multipart form: %w", err)
	}

	// Получаем файл из формы
	file, handler, err := r.FormFile("file")
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error retrieving file from form-data: %w", err)
	}
	defer file.Close()

	// Читаем содержимое файла в память
	fileData, err := io.ReadAll(file)
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error reading file: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	// Формируем DTO с содержимым файла
	saveFileDTO := SaveFileDTO{
		FileName: handler.Filename,
		FileData: fileData,
		UserID:   userID,
	}

	return saveFileDTO, nil
}
