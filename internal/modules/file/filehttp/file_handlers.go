package filehttp

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/file/filedto/request"
	"gophKeeper/internal/modules/file/fileservices"
	"path/filepath"

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
	saveFileDTO, err := request.SaveFileDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("error build DTO for save file: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = p.fileService.SaveFile(r.Context(), saveFileDTO)
	if err != nil {
		logger.Log.Errorf("error saving file: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (p FileHandlers) GetFile(w http.ResponseWriter, r *http.Request) {
	getFileDTO, err := request.GetFileDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	tempFilePath, err := p.fileService.GetFile(r.Context(), getFileDTO)
	if err != nil {
		logger.Log.Errorf("error getting file: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки для скачивания файла
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=\"%s\"",
			filepath.Base(string(tempFilePath))))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(tempFilePath)))

	// Отправляем временный файл в ответ
	http.ServeFile(w, r, tempFilePath)
}

func (p FileHandlers) DeleteFile(w http.ResponseWriter, r *http.Request) {
	deletePwdDTO, err := request.DeleteFileDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.fileService.DeleteFile(r.Context(), deletePwdDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (p FileHandlers) GetAllFiles(w http.ResponseWriter, r *http.Request) {
	allPwdDTO, err := request.AllFileDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	allPasswords, err := p.fileService.GetAllFiles(r.Context(), allPwdDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	marshaledAllPasswords, err := json.Marshal(allPasswords)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledAllPasswords)
}
