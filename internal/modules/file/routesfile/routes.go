package routesfile

import (
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/file/filehttp"
	"gophKeeper/internal/modules/file/fileservices"
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
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.DirSavedFiles).SaveFile(w, req)
	})
	// Получить все пароли
	r.Post("/file/all", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.DirSavedFiles).GetAllFiles(w, req)
	})
	// Удалить пароль
	r.Delete("/file/delete", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.DirSavedFiles).DeleteFile(w, req)
	})
	// Получить конкретный пароль
	r.Get("/file/get/{file_id}", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.DirSavedFiles).GetFile(w, req)
	})
}

func getFilesHandlers(pool *pgxpool.Pool, cryptoKey [32]byte, dirSavedFiles string) *filehttp.FileHandlers {
	fileService := fileservices.NewFileService(pool, cryptoKey, dirSavedFiles)
	pwdHandlers := filehttp.NewFileHandlersHTTP(fileService)
	return &pwdHandlers
}
