package pwd_services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/request"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/response"
	"gophKeeper/pkg/crypto"
)

type PwdService struct {
	pool      *pgxpool.Pool
	cryptoKey [32]byte
}

func NewPwdService(pool *pgxpool.Pool) *PwdService {
	return &PwdService{pool: pool}
}

func (pwd *PwdService) SavePassword(ctx context.Context, dto request.SavePwdDTO) error {
	sql := `INSERT INTO passwords (user_id, title, description, credentials) VALUES ($1, $2, $3)`
	key := [32]byte{}
	encryptedCredentials, err := crypto.Encrypt(key, []byte(dto.Credentials))
	if err != nil {
		return fmt.Errorf("error while encrypting credentials: %w", err)
	}
	_, err = pwd.pool.Exec(ctx, sql, dto.UserID, dto.Title, dto.Description, encryptedCredentials)
	if err != nil {
		return fmt.Errorf("error save password from pwd service: %w", err)
	}

	return nil
}

func (pwd *PwdService) DeletePassword(ctx context.Context, dto request.DeletePwdDTO) error {
	sql := `DELETE FROM passwords WHERE user_id = $1 AND resource = $2;`
	_, err := pwd.pool.Exec(ctx, sql, dto.UserID)
	if err != nil {
		return fmt.Errorf("error save password from pwd service: %w", err)
	}
	return nil
}

func (pwd *PwdService) GetPassword(ctx context.Context, dto request.GetPwdDTO) (response.PwdDTO, error) {
	sql := `SELECT credentials FROM passwords WHERE user_id = $1 AND id = $2;`
	row := pwd.pool.QueryRow(ctx, sql, dto.UserID, dto.PwdID)

	var credentialsData []byte
	err := row.Scan(&credentialsData)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Обработка случая, когда запись не найдена
			return response.PwdDTO{}, fmt.Errorf(
				"password not found for user_id %v and pwd_id %v: %w",
				dto.UserID,
				dto.PwdID,
				err,
			)
		}
		// Обработка других ошибок
		return response.PwdDTO{}, fmt.Errorf("error scanning password from pwd service: %w", err)
	}

	var pwdDTO response.PwdDTO
	if err := json.Unmarshal(credentialsData, &pwdDTO); err != nil {
		return response.PwdDTO{}, fmt.Errorf("error unmarshalling credentials: %w", err)
	}

	return pwdDTO, nil
}

func (pwd *PwdService) GetAllPasswords(ctx context.Context, dto request.AllPwdDTO) ([]InfoByPassword, error) {
	sql := `SELECT resource, login, password FROM passwords WHERE user_id = $1`
	rows, err := pwd.pool.Query(ctx, sql, dto.UserID)
	if err != nil {
		return []InfoByPassword{}, fmt.Errorf("error query get all passwords: %w", err)
	}

	var listPasswords []InfoByPassword

	for rows.Next() {
		var resource string
		var login string
		var password string
		err := rows.Scan(&resource, &login, &password)
		if err != nil {
			return []InfoByPassword{}, fmt.Errorf("error scan get all passwords from pwd service: %w", err)
		}

		listPasswords = append(listPasswords, InfoByPassword{
			Resource: resource,
			Login:    login,
			Password: password,
		})
	}

	return listPasswords, nil
}
