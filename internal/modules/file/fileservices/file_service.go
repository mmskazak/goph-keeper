package fileservices

import (
	"context"
	"fmt"
	"goph-keeper/internal/modules/file/filedto"
	"goph-keeper/pkg/crypto"

	"go.uber.org/zap"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FileService struct {
	pool        *pgxpool.Pool
	zapLogger   *zap.SugaredLogger
	cryptoKey   [32]byte
	maxFileSize int // Максимально допустимый размер файла в байтах

}

func NewFileService(pool *pgxpool.Pool, enKey [32]byte, maxFileSize int, zapLogger *zap.SugaredLogger) *FileService {
	return &FileService{
		pool:        pool,
		cryptoKey:   enKey,
		maxFileSize: maxFileSize,
		zapLogger:   zapLogger,
	}
}

// SaveFile сохраняет файл непосредственно в базу данных.
func (fs *FileService) SaveFile(ctx context.Context, dto filedto.SaveFileDTO) error {
	// Проверяем размер файла
	if (len(dto.FileData)) > fs.maxFileSize {
		return fmt.Errorf("file size exceeds the allowed limit of %d bytes", fs.maxFileSize)
	}

	// Шифруем данные файла
	encryptedFile, err := crypto.Encrypt(fs.cryptoKey, dto.FileData)
	if err != nil {
		return fmt.Errorf("failed to encrypt file data: %w", err)
	}

	// Сохраняем файл и его метаданные в базу данных
	_, err = fs.pool.Exec(ctx, `
		INSERT INTO files (user_id, name_file, file_data) 
		VALUES ($1, $2, $3)`,
		dto.UserID, dto.NameFile, encryptedFile)
	if err != nil {
		return fmt.Errorf("failed to insert file into database: %w", err)
	}

	return nil
}

// GetFile возвращает расшифрованные данные файла.
func (fs *FileService) GetFile(ctx context.Context, dto filedto.GetFileDTO) (
	[]byte,
	string,
	error,
) {
	var nameFile string
	var encryptedData []byte
	// Извлекаем зашифрованные данные из базы данных
	err := fs.pool.QueryRow(ctx, `
		SELECT name_file, file_data FROM files WHERE id = $1 AND user_id = $2`,
		dto.FileID, dto.UserID).Scan(&nameFile, &encryptedData)
	if err != nil {
		fs.zapLogger.Errorf("error retrieving file data for file %d: %v", dto.FileID, err)
		return nil, "", fmt.Errorf("error retrieving file data: %w", err)
	}

	// Расшифровываем данные файла
	decryptedData, err := crypto.Decrypt(fs.cryptoKey, string(encryptedData))
	if err != nil {
		fs.zapLogger.Errorf("error decrypting file %d: %v", dto.FileID, err)
		return nil, "", fmt.Errorf("error decrypting file: %w", err)
	}

	return decryptedData, nameFile, nil
}

// DeleteFile удаляет файл из базы данных.
func (fs *FileService) DeleteFile(ctx context.Context, dto filedto.DeleteFileDTO) error {
	// Удаляем файл из базы данных
	_, err := fs.pool.Exec(ctx, `
		DELETE FROM files WHERE id = $1 AND user_id = $2`,
		dto.FileID, dto.UserID)
	if err != nil {
		fs.zapLogger.Errorf("failed to delete file %d: %v", dto.FileID, err)
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// GetAllFiles возвращает список всех файлов пользователя с их информацией.
func (fs *FileService) GetAllFiles(ctx context.Context, dto filedto.AllFilesDTO) ([]FileInfo, error) {
	rows, err := fs.pool.Query(ctx, `
		SELECT id, name_file FROM files WHERE user_id = $1`,
		dto.UserID)
	if err != nil {
		fs.zapLogger.Errorf("error querying all files: %v", err)
		return nil, fmt.Errorf("error getting all files: %w", err)
	}
	defer rows.Close()

	var files []FileInfo
	for rows.Next() {
		var fileInfo FileInfo
		if err = rows.Scan(&fileInfo.FileID, &fileInfo.NameFile); err != nil {
			fs.zapLogger.Errorf("error scanning row: %v", err)
			return nil, fmt.Errorf("error getting all files: %w", err)
		}
		files = append(files, fileInfo)
	}

	return files, nil
}
