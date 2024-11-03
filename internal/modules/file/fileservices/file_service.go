package fileservices

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"goph-keeper/internal/logger"
	"goph-keeper/internal/modules/file/filedto"
	"goph-keeper/pkg/crypto"
)

type FileService struct {
	pool      *pgxpool.Pool
	cryptoKey [32]byte
}

func NewFileService(pool *pgxpool.Pool, enKey [32]byte) *FileService {
	return &FileService{
		pool:      pool,
		cryptoKey: enKey,
	}
}

// SaveFile сохраняет файл непосредственно в базу данных.
func (fs *FileService) SaveFile(ctx context.Context, dto filedto.SaveFileDTO) error {
	// Шифруем данные файла
	encryptedFile, err := crypto.Encrypt(fs.cryptoKey, dto.FileData)
	if err != nil {
		return fmt.Errorf("failed to encrypt file data: %w", err)
	}

	// Сохраняем файл и его метаданные в базу данных
	_, err = fs.pool.Exec(ctx, `
		INSERT INTO files (user_id, title, description, file_data) 
		VALUES ($1, $2, $3, $4)`,
		dto.UserID, dto.Title, dto.Description, encryptedFile)
	if err != nil {
		return fmt.Errorf("failed to insert file into database: %w", err)
	}

	return nil
}

// GetFile возвращает расшифрованные данные файла.
func (fs *FileService) GetFile(ctx context.Context, dto filedto.GetFileDTO) ([]byte, error) {
	var encryptedData []byte

	// Извлекаем зашифрованные данные из базы данных
	err := fs.pool.QueryRow(ctx, `
		SELECT file_data FROM files WHERE id = $1 AND user_id = $2`,
		dto.FileID, dto.UserID).Scan(&encryptedData)
	if err != nil {
		logger.Log.Errorf("error retrieving file data for file %d: %v", dto.FileID, err)
		return nil, fmt.Errorf("error retrieving file data: %w", err)
	}

	// Расшифровываем данные файла
	decryptedData, err := crypto.Decrypt(fs.cryptoKey, string(encryptedData))
	if err != nil {
		logger.Log.Errorf("error decrypting file %d: %v", dto.FileID, err)
		return nil, fmt.Errorf("error decrypting file: %w", err)
	}

	return decryptedData, nil
}

// DeleteFile удаляет файл из базы данных.
func (fs *FileService) DeleteFile(ctx context.Context, dto filedto.DeleteFileDTO) error {
	// Удаляем файл из базы данных
	_, err := fs.pool.Exec(ctx, `
		DELETE FROM files WHERE id = $1 AND user_id = $2`,
		dto.FileID, dto.UserID)
	if err != nil {
		logger.Log.Errorf("failed to delete file %d: %v", dto.FileID, err)
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// GetAllFiles возвращает список всех файлов пользователя с их информацией.
func (fs *FileService) GetAllFiles(ctx context.Context, dto filedto.AllFilesDTO) ([]FileInfo, error) {
	rows, err := fs.pool.Query(ctx, `
		SELECT id, title, description FROM files WHERE user_id = $1`,
		dto.UserID)
	if err != nil {
		logger.Log.Errorf("error querying all files: %v", err)
		return nil, fmt.Errorf("error getting all files: %w", err)
	}
	defer rows.Close()

	var files []FileInfo
	for rows.Next() {
		var fileInfo FileInfo
		if err = rows.Scan(&fileInfo.FileID, &fileInfo.Title, &fileInfo.Description); err != nil {
			logger.Log.Errorf("error scanning row: %v", err)
			return nil, fmt.Errorf("error getting all files: %w", err)
		}
		files = append(files, fileInfo)
	}

	return files, nil
}
