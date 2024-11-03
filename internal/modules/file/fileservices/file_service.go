package fileservices

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"goph-keeper/internal/logger"
	"goph-keeper/internal/modules/file/filedto"
	"goph-keeper/pkg/crypto"
	"os"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FileService struct {
	pool          *pgxpool.Pool
	dirSavedFiles string // Путь к директории для хранения файлов
	cryptoKey     [32]byte
}

func NewFileService(pool *pgxpool.Pool, enKey [32]byte, dirSavedFiles string) *FileService {
	return &FileService{
		pool:          pool,
		cryptoKey:     enKey,
		dirSavedFiles: dirSavedFiles,
	}
}

// SaveFile сохраняет файл на сервере и сохраняет метаданные в базу данных.
func (fs *FileService) SaveFile(ctx context.Context, dto filedto.SaveFileDTO) error {
	randomFileName, err := generateRandomFileName()
	if err != nil {
		return err
	}

	destPath := filepath.Join(fs.dirSavedFiles, randomFileName)
	logger.Log.Infoln(destPath)

	// Создаем файл в целевой директории
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("create file.proto %s failed", destPath)
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			logger.Log.Errorf("Failed to close file.proto %s, err: %s", destPath, err)
		}
	}(destFile)

	// Шифруем данные
	encryptedFile, err := crypto.Encrypt(fs.cryptoKey, dto.FileData)
	if err != nil {
		return fmt.Errorf("encrypt file.proto %s failed", destPath)
	}

	// Пишем содержимое файла в созданный файл
	if _, err = destFile.Write([]byte(encryptedFile)); err != nil {
		return fmt.Errorf("write file.proto %s failed", destPath)
	}

	// Сохраняем информацию о файле в базе данных
	_, err = fs.pool.Exec(ctx, "INSERT INTO files (user_id, title, description, path_to_file) VALUES ($1, $2, $3, $4)",
		dto.UserID, dto.Title, dto.Description, destPath)
	if err != nil {
		return fmt.Errorf("insert query  %s failed", destPath)
	}

	return fmt.Errorf("save file.proto %s failed", destPath)
}

// DeleteFile удаляет файл с сервера и удаляет метаданные из базы данных.
func (fs *FileService) DeleteFile(ctx context.Context, dto filedto.DeleteFileDTO) error {
	var filePath string
	// Получаем путь к файлу и проверяем, что файл принадлежит пользователю
	err := fs.pool.QueryRow(ctx, "SELECT path_to_file FROM files WHERE id = $1 AND user_id = $2",
		dto.FileID, dto.UserID).
		Scan(&filePath)
	if err != nil {
		return fmt.Errorf("select query file.proto %s failed", dto.FileID)
	}

	// Удаляем файл с диска
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("delete file.proto %s failed", filePath)
	}

	// Удаляем информацию о файле из базы данных
	_, err = fs.pool.Exec(ctx, "DELETE FROM files WHERE id = $1 AND user_id = $2", dto.FileID, dto.UserID)
	if err != nil {
		return fmt.Errorf("delete query file.proto %s failed", dto.FileID)
	}
	return nil
}

// GetFile возвращает путь к временно созданному расшифрованному файлу.
func (fs *FileService) GetFile(ctx context.Context, dto filedto.GetFileDTO) (string, error) {
	var filePath string
	// Получаем путь к файлу и проверяем, что файл принадлежит пользователю
	err := fs.pool.QueryRow(ctx, "SELECT path_to_file FROM files WHERE id = $1 AND user_id = $2",
		dto.FileID, dto.UserID).
		Scan(&filePath)
	if err != nil {
		logger.Log.Errorf("get file.proto %s error: %v", dto.FileID, err)
		return "", fmt.Errorf("error retrieving file.proto path: %w", err)
	}

	// Читаем зашифрованный файл с диска
	encryptedData, err := os.ReadFile(filePath)
	if err != nil {
		logger.Log.Errorf("get file.proto %s error: %v", dto.FileID, err)
		return "", fmt.Errorf("error reading file.proto: %w", err)
	}

	// Преобразуем зашифрованные данные в строку
	encryptedString := string(encryptedData)

	// Расшифровываем данные
	decryptedData, err := crypto.Decrypt(fs.cryptoKey, encryptedString) // Передаем строку
	if err != nil {
		logger.Log.Errorf("decrypt file.proto %s error: %v", dto.FileID, err)
		return "", fmt.Errorf("error decrypting file.proto: %w", err)
	}

	// Создаем временный файл для расшифрованных данных
	tempFile, err := os.CreateTemp("", "decrypted_*.tmp")
	if err != nil {
		return "", fmt.Errorf("error creating temp file.proto: %w", err)
	}
	defer func() {
		time.AfterFunc(5*time.Second, func() { //nolint:gomnd //5 секунд
			if err = os.Remove(tempFile.Name()); err != nil {
				logger.Log.Errorf("failed to remove temp file.proto: %v", err)
			}
		})
	}()

	// Записываем расшифрованные данные во временный файл
	if _, err = tempFile.Write(decryptedData); err != nil {
		return "", fmt.Errorf("error writing to temp file.proto: %w", err)
	}

	// Закрываем файл
	if err := tempFile.Close(); err != nil {
		return "", fmt.Errorf("error closing temp file.proto: %w", err)
	}

	// Возвращаем путь к временному файлу
	return tempFile.Name(), nil
}

// GetAllFiles возвращает список всех файлов пользователя с их информацией.
func (fs *FileService) GetAllFiles(ctx context.Context, dto filedto.AllFilesDTO) ([]FileInfo, error) {
	rows, err := fs.pool.Query(ctx, "SELECT id, title, description FROM files WHERE user_id = $1", dto.UserID)
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

func generateRandomFileName() (string, error) {
	b := make([]byte, 16) //nolint:gomnd //16 байт = 128 бит случайных данных
	if _, err := rand.Read(b); err != nil {
		logger.Log.Errorf("error generating random filename: %v", err)
		return "", fmt.Errorf("error generating random filename: %w", err)
	}
	return hex.EncodeToString(b), nil
}
