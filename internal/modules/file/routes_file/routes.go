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
		getFilesHandlers(pool, cfg.DirSavedFiles).SaveFile(w, req)
	})
	//Получить все пароли
	r.Post("/file/all", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.DirSavedFiles).GetAllFiles(w, req)
	})
	//Удалить пароль
	r.Delete("/file/delete", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.DirSavedFiles).DeleteFile(w, req)
	})
	//Получить конкретный пароль
	r.Post("/file/get", func(w http.ResponseWriter, req *http.Request) {
		getFilesHandlers(pool, cfg.DirSavedFiles).GetFile(w, req)
	})

	return r
}

func getFilesHandlers(pool *pgxpool.Pool, dirSavedFiles string) *file_http.FileHandlers {
	fileService := file_services.NewFileService(pool, dirSavedFiles)
	pwdHandlers := file_http.NewFileHandlersHTTP(fileService)
	return &pwdHandlers
}
