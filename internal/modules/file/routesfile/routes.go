package routesfile

import (
	"goph-keeper/internal/config"
	"goph-keeper/internal/modules/file/filehttp"
	"goph-keeper/internal/modules/file/fileservices"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegistrationRoutesFile(
	r chi.Router,
	pool *pgxpool.Pool,
	cfg *config.Config,
) {
	// Сохранить пароль
	r.Post("/file/save", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.MaxFileSize).SaveFile(w, req)
	})
	// Получить все пароли
	r.Get("/file/all", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.MaxFileSize).GetAllFiles(w, req)
	})
	// Удалить пароль
	r.Get("/file/delete", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.MaxFileSize).DeleteFile(w, req)
	})
	// Получить конкретный пароль
	r.Get("/file/get/{file_id}", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.MaxFileSize).GetFile(w, req)
	})
}

func getFilesHandlers(pool *pgxpool.Pool, cryptoKey [32]byte, maxFileSize int) *filehttp.FileHandlers {
	fileService := fileservices.NewFileService(pool, cryptoKey, maxFileSize)
	pwdHandlers := filehttp.NewFileHandlersHTTP(fileService)
	return &pwdHandlers
}
