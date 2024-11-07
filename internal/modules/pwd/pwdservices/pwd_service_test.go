package pwdservices

import (
	"context"
	"errors"
	"goph-keeper/internal/modules/pwd/pwddto"
	"goph-keeper/internal/modules/pwd/valueobj"
	"goph-keeper/internal/storage/mocks"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPwdService_SavePassword(t *testing.T) {
	mockPool := new(mocks.MockDatabase)

	ctx := context.Background()
	dto := pwddto.SavePwdDTO{
		UserID: 1,
		Title:  "test title",
		Credentials: valueobj.Credentials{
			Login:    "admin",
			Password: "12345",
		},
	}
	// Строка длиной 32 символа
	strKey := "MySecretEncryptionKey1234567890a"
	// Преобразуем строку в массив байтов
	var key [32]byte
	copy(key[:], strKey)

	ct := pgconn.NewCommandTag("success")
	mockPool.On("Exec", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(ct, nil)
	s := NewPwdService(mockPool, key, zap.NewNop().Sugar())
	err := s.SavePassword(ctx, &dto)
	require.Nil(t, err)
}

func TestPwdService_DeletePassword(t *testing.T) {
	mockPool := new(mocks.MockDatabase)
	mockRow := new(mocks.MockRow)
	ctx := context.Background()
	dto := pwddto.DeletePwdDTO{
		UserID: 1,
		PwdID:  "1",
	}

	// Строка длиной 32 символа
	strKey := "MySecretEncryptionKey1234567890a"
	// Преобразуем строку в массив байтов
	var key [32]byte
	copy(key[:], strKey)

	mockRow.On("Scan", mock.Anything).Return(nil)
	mockPool.On("QueryRow", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(mockRow, nil)

	s := NewPwdService(mockPool, key, zap.NewNop().Sugar())
	err := s.DeletePassword(ctx, &dto)
	require.Nil(t, err)
}

func TestPwdService_GetPassword(t *testing.T) {
	mockPool := new(mocks.MockDatabase)
	ctx := context.Background()
	dto := pwddto.GetPwdDTO{
		UserID: 1,
		PwdID:  "1",
	}

	// Строка длиной 32 символа
	strKey := "MySecretEncryptionKey1234567890a"
	// Преобразуем строку в массив байтов
	var key [32]byte
	copy(key[:], strKey)
	mkRow := new(mocks.MockRow)

	errDecrypt := errors.New("error decrypt")
	mockPool.On("QueryRow", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(mkRow)

	mkRow.On("Scan", mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			if dest1, ok := args.Get(0).(*string); ok {
				*dest1 = "1"
			}
			if dest2, ok := args.Get(1).(*string); ok {
				*dest2 = "title"
			}
			if dest3, ok := args.Get(2).(*[]byte); ok {
				*dest3 = []byte("credentials")
			}
		}).
		Return(errDecrypt)

	s := NewPwdService(mockPool, key, zap.NewNop().Sugar())
	_, err := s.GetPassword(ctx, &dto)
	assert.EqualError(t,
		err,
		"error scanning password from pwd service: error MockRow func Scan: error decrypt")
}

func TestPwdService_GetAllPasswords(t *testing.T) {
	ctx := context.Background()
	mockPool := new(mocks.MockDatabase)
	mockRows := new(mocks.MockRows)
	// Строка длиной 32 символа
	strKey := "MySecretEncryptionKey1234567890a"
	// Преобразуем строку в массив байтов
	var key [32]byte
	copy(key[:], strKey)
	s := NewPwdService(mockPool, key, zap.NewNop().Sugar())
	dto := pwddto.AllPwdDTO{
		UserID: 1,
	}
	queryError := errors.New("query error")
	mockPool.On("Query", ctx,
		mock.Anything,
		mock.Anything,
	).Return(mockRows, queryError)

	_, err := s.GetAllPasswords(ctx, &dto)
	assert.ErrorIs(t, err, queryError)
}
