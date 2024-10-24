package request

import (
	"fmt"
	"gophKeeper/internal/helpers"
	"io"
	"net/http"
)

type SaveFileDTO struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description []byte `json:"description"` // Описание в байтовом формате
	FileData    []byte `json:"file_data"`   // Содержимое файла
}

func SaveFileDTOFromHTTP(r *http.Request) (SaveFileDTO, error) {
	// Ограничиваем размер загружаемого файла до 10MB
	const maxUploadSize = 10 << 20
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return SaveFileDTO{}, fmt.Errorf("error parsing multipart form: %w", err)
	}

	// Получаем файл из формы
	file, _, err := r.FormFile("file")
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error retrieving file from form-data: %w", err)
	}
	defer file.Close()

	// Читаем содержимое файла
	fileData, err := io.ReadAll(file)
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error reading file: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error retrieving userID from context: %w", err)
	}

	// Читаем JSON-данные из формы
	title := r.FormValue("title")
	description := r.FormValue("description")

	// Формируем DTO с данными
	saveFileDTO := SaveFileDTO{
		UserID:      userID,
		Title:       title,
		Description: []byte(description),
		FileData:    fileData,
	}

	return saveFileDTO, nil
}
