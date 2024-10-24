package pwd_services

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gophKeeper/internal/modules/pwd/pwd_dto/request"
	"gophKeeper/internal/modules/pwd/value_obj"
	"gophKeeper/internal/storage/mocks"
	"testing"
)

func TestPwdService_SavePassword(t *testing.T) {

	mockPool := new(mocks.MockDatabase)

	ctx := context.Background()
	dto := request.SavePwdDTO{
		UserID:      1,
		Title:       "test title",
		Description: "test description",
		Credentials: value_obj.Credentials{
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
	err := s.SavePassword(ctx, dto)
	require.Nil(t, err)

}
