package authservice

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	dto2 "goph-keeper/internal/modules/auth/authdto"
	"goph-keeper/internal/modules/auth/authservices/authhashpwd"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthService struct {
	pool      *pgxpool.Pool
	zapLogger *zap.SugaredLogger
}

func NewAuthService(pool *pgxpool.Pool, zapLogger *zap.SugaredLogger) *AuthService {
	return &AuthService{
		pool:      pool,
		zapLogger: zapLogger,
	}
}

// Registration выполняет регистрацию нового пользователя.
func (a *AuthService) Registration(ctx context.Context, regDTO *dto2.RegistrationDTO) (int, error) {
	// Генерация хэша пароля
	hashedPassword, err := authhashpwd.GenerateHashFromPassword(regDTO.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	// SQL-запрос для добавления пользователя
	sql := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;"
	row := a.pool.QueryRow(ctx, sql, regDTO.Username, hashedPassword)

	// Сканируем полученный id
	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error while inserting user: %w", err)
	}

	return id, nil
}

// Login выполняет проверку логина и пароля пользователя.
func (a *AuthService) Login(ctx context.Context, logDTO *dto2.LoginDTO) (int, error) {
	// SQL-запрос для получения пользователя по имени
	sql := "SELECT id, password FROM users WHERE username = $1;"
	row := a.pool.QueryRow(ctx, sql, logDTO.Username)

	// Проверка на наличие пользователя
	var id int
	var hashedPassword string
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return 0, fmt.Errorf("error while scanning user: %w", err)
	}

	// Сравнение пароля с хранимым хэшем
	isLogin := authhashpwd.IsEqualHashedPassword(hashedPassword, logDTO.Password)
	if !isLogin {
		return 0, errors.New("password incorrect")
	}

	return id, nil
}
