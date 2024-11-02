package filehttp

import (
	"encoding/json"
	"fmt"
	"goph-keeper/internal/logger"
	"goph-keeper/internal/modules/file/filedto"
	"goph-keeper/internal/modules/file/fileservices"
	"path/filepath"
	"strconv"

	"net/http"
)

type FileHandlers struct {
	fileService fileservices.IFileService
}

func NewFileHandlersHTTP(service fileservices.IFileService) FileHandlers {
	return FileHandlers{
		fileService: service,
	}
}

func (p FileHandlers) SaveFile(w http.ResponseWriter, r *http.Request) {
	saveFileDTO, err := filedto.SaveFileDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("error build DTO for save file.proto: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = p.fileService.SaveFile(r.Context(), saveFileDTO)
	if err != nil {
		logger.Log.Errorf("error saving file.proto: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		logger.Log.Errorf("error writing response: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p FileHandlers) GetFile(w http.ResponseWriter, r *http.Request) {
	getFileDTO, err := filedto.GetFileDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("error build DTO for get file.proto: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	tempFilePath, err := p.fileService.GetFile(r.Context(), getFileDTO)
	if err != nil {
		logger.Log.Errorf("error getting file.proto: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки для скачивания файла
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=\"%q\"", filepath.Base(tempFilePath)))
	w.Header().Set("Content-Length", strconv.Itoa(len(tempFilePath)))

	// Отправляем временный файл в ответ
	http.ServeFile(w, r, tempFilePath)
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
