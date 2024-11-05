package pwdservices

import (
	"context"
	"fmt"
	"goph-keeper/internal/modules/pwd/pwddto"
	"goph-keeper/internal/modules/pwd/valueobj"
	"goph-keeper/internal/storage/mocks"
	"testing"

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
	s := NewPwdService(mockPool, key)
	err := s.SavePassword(ctx, &dto)
	require.Nil(t, err)
}

func TestPwdService_DeletePassword(t *testing.T) {
	mockPool := new(mocks.MockDatabase)
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
	ct := pgconn.NewCommandTag("success")
	mockPool.On("Exec", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(ct, nil)
	s := NewPwdService(mockPool, key)
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

	mockPool.On("QueryRow", ctx,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(mkRow)

	mkRow.On("Scan", mock.Anything).
		Run(func(args mock.Arguments) {
			if dest, ok := args.Get(0).(*[]byte); ok {
				*dest = []byte("error credentials")
			}
		}).
		Return(nil)

	s := NewPwdService(mockPool, key)
	pwd, err := s.GetPassword(ctx, &dto)
	assert.EqualError(t, err, "error decrypt for GetPassword invalid format: expected nonce:ciphertext")
	fmt.Println(pwd)
}
