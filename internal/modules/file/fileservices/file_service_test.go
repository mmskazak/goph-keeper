package fileservices

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"goph-keeper/internal/modules/file/filedto"
	mocksFile "goph-keeper/internal/modules/file/mocks"
	"testing"
)

func TestFileService_SaveFile(t *testing.T) {
	ctx := context.Background()
	mockPool := new(mocksFile.MockDatabase)
	maxFileSize := 1 * (1024 * 1024) // 1Mb
	// Строка длиной 32 символа
	strKey := "MySecretEncryptionKey1234567890a"
	// Преобразуем строку в массив байтов
	var key [32]byte
	copy(key[:], strKey)
	fs := NewFileService(mockPool, key, maxFileSize, zap.NewNop().Sugar())
	saveFileDTO := filedto.SaveFileDTO{
		UserID:   1,
		NameFile: "test.txt",
		FileData: []byte("test"),
	}
	commandTag := pgconn.NewCommandTag("done")
	mockPool.On("Exec", ctx, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(commandTag, nil)
	err := fs.SaveFile(ctx, saveFileDTO)
	assert.NoError(t, err)
}
