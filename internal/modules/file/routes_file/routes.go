package routes_file

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/file/file_http"
	"gophKeeper/internal/modules/file/file_services"
	"net/http"
)

func RegistrationRoutesFile(
	r chi.Router,
	pool *pgxpool.Pool,
	cfg *config.Config,
) chi.Router {
	//Сохранить пароль
	r.Post("/file/save", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.DirSavedFiles).SaveFile(w, req)
	})
	//Получить все пароли
	r.Post("/file/all", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.DirSavedFiles).GetAllFiles(w, req)
	})
	//Удалить пароль
	r.Delete("/file/delete", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.DirSavedFiles).DeleteFile(w, req)
	})
	//Получить конкретный пароль
	r.Get("/file/get/{file_id}", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.EncryptionKey, cfg.DirSavedFiles).GetFile(w, req)
	})

	return r
}

func getFilesHandlers(pool *pgxpool.Pool, cryptoKey [32]byte, dirSavedFiles string) *file_http.FileHandlers {
	fileService := file_services.NewFileService(pool, cryptoKey, dirSavedFiles)
	pwdHandlers := file_http.NewFileHandlersHTTP(fileService)
	return &pwdHandlers
}
