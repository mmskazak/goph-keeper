package file_services

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/modules/file/file_dto/request"
	"io"
	"os"
	"path/filepath"
)

type FileService struct {
	pool          *pgxpool.Pool
	dirSavedFiles string // Путь к директории для хранения файлов
}

func NewFileService(pool *pgxpool.Pool, dirSavedFiles string) *FileService {
	return &FileService{
		pool:          pool,
		dirSavedFiles: dirSavedFiles,
	}
}

// SaveFile сохраняет файл на сервере и сохраняет метаданные в базу данных
func (fs *FileService) SaveFile(ctx context.Context, dto request.SaveFileDTO) error {
	destPath := filepath.Join(fs.dirSavedFiles, dto.FileName)

	// Открываем исходный файл
	srcFile, err := os.Open(dto.FilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Создаем файл в целевой директории
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Копируем содержимое файла
	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}

	// Сохраняем информацию о файле в базе данных
	_, err = fs.pool.Exec(ctx, "INSERT INTO files (user_id, file_name, file_path) VALUES ($1, $2, $3)",
		dto.UserID, dto.FileName, destPath)

	return err
}

// DeleteFile удаляет файл с сервера и удаляет метаданные из базы данных
func (fs *FileService) DeleteFile(ctx context.Context, dto request.DeleteFileDTO) error {
	var filePath string

	// Получаем путь к файлу и проверяем, что файл принадлежит пользователю
	err := fs.pool.QueryRow(ctx, "SELECT file_path FROM files WHERE id = $1 AND user_id = $2", dto.FileID, dto.UserID).Scan(&filePath)
	if err != nil {
		return err
	}

	// Удаляем файл с диска
	if err := os.Remove(filePath); err != nil {
		return err
	}

	// Удаляем информацию о файле из базы данных
	_, err = fs.pool.Exec(ctx, "DELETE FROM files WHERE id = $1 AND user_id = $2", dto.FileID, dto.UserID)
	return err
}

// GetFile возвращает путь к файлу на сервере
func (fs *FileService) GetFile(ctx context.Context, dto request.GetFileDTO) (string, error) {
	var filePath string

	// Получаем путь к файлу и проверяем, что файл принадлежит пользователю
	err := fs.pool.QueryRow(ctx, "SELECT file_path FROM files WHERE id = $1 AND user_id = $2", dto.FileID, dto.UserID).Scan(&filePath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// GetAllFiles возвращает список всех файлов пользователя с их информацией
func (fs *FileService) GetAllFiles(ctx context.Context, dto request.AllFilesDTO) ([]FileInfo, error) {
	rows, err := fs.pool.Query(ctx, "SELECT file_name, file_path FROM files WHERE user_id = $1", dto.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []FileInfo
	for rows.Next() {
		var fileInfo FileInfo
		if err := rows.Scan(&fileInfo.Name, &fileInfo.PathToFile); err != nil {
			return nil, err
		}
		files = append(files, fileInfo)
	}

	return files, nil
}
