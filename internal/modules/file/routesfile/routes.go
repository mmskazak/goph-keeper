package routesfile

import (
	"goph-keeper/internal/modules/file/filehttp"
	"goph-keeper/internal/modules/file/fileservices"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
)

func RegistrationRoutesFile(
	r chi.Router,
	fileService *fileservices.FileService,
	zapLogger *zap.SugaredLogger,
) {
	// Сохранить пароль
	r.Post("/file/save", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(fileService, zapLogger).SaveFile(w, req)
	})
	// Получить все пароли
	r.Get("/file/all", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(fileService, zapLogger).GetAllFiles(w, req)
	})
	// Удалить пароль
	r.Get("/file/delete", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(fileService, zapLogger).DeleteFile(w, req)
	})
	// Получить конкретный пароль
	r.Get("/file/get/{file_id}", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(fileService, zapLogger).GetFile(w, req)
	})
}

func getFilesHandlers(fileService *fileservices.FileService, zapLogger *zap.SugaredLogger) *filehttp.FileHandlers {
	pwdHandlers := filehttp.NewFileHandlersHTTP(fileService, zapLogger)
	return &pwdHandlers
}
