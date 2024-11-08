package pwdservices

import (
	"context"
	"errors"
	mocksPwd "goph-keeper/internal/modules/pwd/mocks"
	"goph-keeper/internal/modules/pwd/pwddto"
	"goph-keeper/internal/modules/pwd/valueobj"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPwdService_SavePassword(t *testing.T) {
	mockPool := new(mocksPwd.MockDatabase)

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
	mockPool := new(mocksPwd.MockDatabase)
	mockRow := new(mocksPwd.MockRow)
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

func TestPwdService_GetPasswordErrorDecrypt(t *testing.T) {
	mockPool := new(mocksPwd.MockDatabase)
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
	mkRow := new(mocksPwd.MockRow)

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

func TestPwdService_GetPassword(t *testing.T) {
	mockPool := new(mocksPwd.MockDatabase)
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
	mkRow := new(mocksPwd.MockRow)

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
				*dest3 = []byte("{\"login\": \"myemail@example.com\", \"password\": \"17c5d99202bd0e1141d3dcee:2ac2a44261925c17213821b8c4\"}")
			}
		}).
		Return(nil)

	s := NewPwdService(mockPool, key, zap.NewNop().Sugar())
	responsePwdDTO, err := s.GetPassword(ctx, &dto)
	require.NoError(t, err)
	expectedDTO := pwddto.ResponsePwdDTO{
		PwdID: "1",
		Title: "title",
		Credentials: valueobj.Credentials{
			Login:    "myemail@example.com",
			Password: "mypassword123",
		},
	}
	assert.Equal(t, expectedDTO, responsePwdDTO)
}

func TestPwdService_GetAllPasswordsQueryError(t *testing.T) {
	ctx := context.Background()
	mockPool := new(mocksPwd.MockDatabase)
	mockRows := new(mocksPwd.MockRows)
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

func TestPwdService_GetAllPasswords(t *testing.T) {
	ctx := context.Background()
	mockPool := new(mocksPwd.MockDatabase)
	mockRows := new(mocksPwd.MockRows)
	// Строка длиной 32 символа
	strKey := "MySecretEncryptionKey1234567890a"
	// Преобразуем строку в массив байтов
	var key [32]byte
	copy(key[:], strKey)
	s := NewPwdService(mockPool, key, zap.NewNop().Sugar())
	dto := pwddto.AllPwdDTO{
		UserID: 1,
	}
	mockPool.On("Query", ctx,
		mock.Anything,
		mock.Anything,
	).Return(mockRows, nil)

	mockRows.On("Next").Return(true).Once()  // Первый вызов вернет true
	mockRows.On("Next").Return(false).Once() // Второй вызов вернет false
	mockRows.On("Err").Return(nil).Once()
	mockRows.On("Scan",
		mock.AnythingOfType("*string"),
		mock.AnythingOfType("*string"),
		mock.AnythingOfType("*[]uint8"),
	).
		Run(func(args mock.Arguments) {
			id, ok := args.Get(0).(*string)
			if ok {
				*id = "1"
			}
			title, ok := args.Get(1).(*string)
			if ok {
				*title = "http://google.com"
			}
			credential, ok := args.Get(2).(*[]byte)
			if ok {
				*credential = []byte("{\"login\": \"myemail@example.com\", \"password\": \"17c5d99202bd0e1141d3dcee:2ac2a44261925c17213821b8c4\"}")
			}
		}).Return(nil)

	listPasswords, err := s.GetAllPasswords(ctx, &dto)
	assert.NoError(t, err)
	assert.Equal(t, "mypassword123", listPasswords[0].Credentials.Password)
}

func TestPwdService_GetAllPasswordsErrUnmarshallingCredentials(t *testing.T) {
	ctx := context.Background()
	mockPool := new(mocksPwd.MockDatabase)
	mockRows := new(mocksPwd.MockRows)
	// Строка длиной 32 символа
	strKey := "MySecretEncryptionKey1234567890a"
	// Преобразуем строку в массив байтов
	var key [32]byte
	copy(key[:], strKey)
	s := NewPwdService(mockPool, key, zap.NewNop().Sugar())
	dto := pwddto.AllPwdDTO{
		UserID: 1,
	}
	mockPool.On("Query", ctx,
		mock.Anything,
		mock.Anything,
	).Return(mockRows, nil)

	mockRows.On("Next").Return(true).Once()  // Первый вызов вернет true
	mockRows.On("Next").Return(false).Once() // Второй вызов вернет false
	mockRows.On("Err").Return(nil).Once()
	errScan := errors.New("error scanning password from pwd service")
	mockRows.On("Scan",
		mock.AnythingOfType("*string"),
		mock.AnythingOfType("*string"),
		mock.AnythingOfType("*[]uint8"),
	).
		Run(func(args mock.Arguments) {
			id, ok := args.Get(0).(*string)
			if ok {
				*id = "1"
			}
			title, ok := args.Get(1).(*string)
			if ok {
				*title = "http://google.com"
			}
			credential, ok := args.Get(2).(*[]byte)
			if ok {
				*credential = []byte("login\": \"myemail@example.com\", \"password\": \"17c5d99202bd0e1141d3dcee:2ac2a44261925c17213821b8c4")
			}
		}).Return(errScan)

	listPasswords, err := s.GetAllPasswords(ctx, &dto)
	assert.ErrorIs(t, err, errScan)
	assert.Equal(t, []pwddto.ResponsePwdDTO{}, listPasswords)
}
