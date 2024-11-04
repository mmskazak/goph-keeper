package filehttp

import (
	"encoding/json"
	"fmt"
	"goph-keeper/internal/logger"
	"goph-keeper/internal/modules/file/filedto"
	"goph-keeper/internal/modules/file/fileservices"
	"strconv"

	"net/http"
)

const ContentType = "Content-Type"

type FileHandlers struct {
	fileService fileservices.IFileService
}

func NewFileHandlersHTTP(service fileservices.IFileService) FileHandlers {
	return FileHandlers{
		fileService: service,
	}
}

func (p FileHandlers) SaveFile(w http.ResponseWriter, r *http.Request) {
	// Преобразуем HTTP-запрос в DTO
	saveFileDTO, err := filedto.SaveFileDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Ошибка при построении DTO для сохранения файла: %v", err)
		w.Header().Set(ContentType, "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid file data",
		})
		return
	}

	// Сохраняем файл
	err = p.fileService.SaveFile(r.Context(), saveFileDTO)
	if err != nil {
		logger.Log.Errorf("Ошибка сохранения файла: %v", err)
		w.Header().Set(ContentType, "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to save file.",
		})
		return
	}

	// Успешный ответ при сохранении файла
	w.Header().Set(ContentType, "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "File saved successfully.",
	})
}

func (p FileHandlers) GetFile(w http.ResponseWriter, r *http.Request) {
	getFileDTO, err := filedto.GetFileDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("error build DTO for get file.proto: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// Получаем байты файла из сервиса
	fileData, err := p.fileService.GetFile(r.Context(), getFileDTO)
	if err != nil {
		logger.Log.Errorf("error getting file.proto: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки для скачивания файла
	w.Header().Set(ContentType, "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%q\"", getFileDTO.FileID))
	w.Header().Set("Content-Length", strconv.Itoa(len(fileData)))

	// Отправляем байты файла в ответ
	if _, err := w.Write(fileData); err != nil {
		logger.Log.Errorf("error writing file data to response: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p FileHandlers) DeleteFile(w http.ResponseWriter, r *http.Request) {
	deletePwdDTO, err := filedto.DeleteFileDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("error build DTO for delete file.proto: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.fileService.DeleteFile(r.Context(), deletePwdDTO)
	if err != nil {
		logger.Log.Errorf("error deleting file.proto: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		logger.Log.Errorf("error writing response: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p FileHandlers) GetAllFiles(w http.ResponseWriter, r *http.Request) {
	allPwdDTO, err := filedto.AllFileDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	allPasswords, err := p.fileService.GetAllFiles(r.Context(), allPwdDTO)
	if err != nil {
		logger.Log.Errorf("error getting all files: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	marshaledAllPasswords, err := json.Marshal(allPasswords)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(marshaledAllPasswords)
	if err != nil {
		logger.Log.Errorf("error writing response: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
