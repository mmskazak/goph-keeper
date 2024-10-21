package file_http

import (
	"encoding/json"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/file/file_dto/request"
	"gophKeeper/internal/modules/file/file_services"

	"net/http"
)

type FileHandlers struct {
	fileService file_services.IFileService
}

func NewFileHandlersHTTP(service file_services.IFileService) FileHandlers {
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
	password, err := p.fileService.GetFile(r.Context(), getFileDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(password))
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
